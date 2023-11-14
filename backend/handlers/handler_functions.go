package handlers

import (
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/tokenPackage"
)

func ConvertToCustomXML(xmlData []byte, singularTableName string) (string, error) {
	// Define the opening and closing tags
	opening_tag := "<" + singularTableName
	closing_tag := "</" + singularTableName

	// Convert XML data to string and replace <map> tags with the opening and closing tags
	xmlString := string(xmlData)
	xmlString = strings.ReplaceAll(xmlString, "<map>", opening_tag+">")
	xmlString = strings.ReplaceAll(xmlString, "</map>", closing_tag+">")

	// Replace the first occurence of the opening tag with a plural opening tag
	xmlString = strings.Replace(xmlString, opening_tag+">", opening_tag+"s>", 1)

	// Replace the last occurence of the closing tag with a plural closing tag
	lastOrgIndex := strings.LastIndex(xmlString, closing_tag+">")
	if lastOrgIndex != -1 {
		xmlString = xmlString[:lastOrgIndex] + closing_tag + "s>" + xmlString[lastOrgIndex+len(closing_tag+">"):]
	}

	return xmlString, nil
}

// CheckEmptyFields checks if any fields in a struct are empty.
func CheckEmptyFields(i interface{}) bool {
	v := reflect.ValueOf(i)

	// Loop through the struct's fields
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String && v.Field(i).String() == "" {
			return true
		}
	}

	return false
}

// ExtractDBAndToken extracts JWT token and database connection from the request context.
func ExtractDBAndToken(c *gin.Context) (*jwt.Token, string, *database.Queries) {
	// Extract JWT token from the request header
	tokenString := tokenPackage.ExtractJWTTokenFromHeader(c.Request)

	// Check if the token is missing or invalid
	if tokenString == "" {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "JWT token is missing or invalid",
		})
		return nil, "", nil
	}

	// Parse and validate the JWT token
	token, err := tokenPackage.ParseAndValidateJWTToken(tokenString)
	if err != nil {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "Invalid JWT token",
		})
		return nil, "", nil
	}

	// Check for revoked token
	for _, tok := range tokenPackage.TokenMap {
		if tok.Claims.(jwt.MapClaims)["ID"] == token.Claims.(jwt.MapClaims)["ID"] {
			token = tok
			break
		}
	}

	// Check if the token is revoked
	if token.Claims.(jwt.MapClaims)["Revoked"] == true {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "Trying to use revoked JWT token",
		})
		return nil, "", nil
	}

	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return nil, "", nil
	}

	return token, tokenString, db
}

// VerifyTeamOwnershipFromParam verifies if the user has ownership of a team based on the team name parameter.
func VerifyTeamOwnershipFromParam(c *gin.Context, token *jwt.Token, db *database.Queries, teamNameParam string) database.Team {
	// Check if the team name is missing
	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Team name is missing",
		})
		return database.Team{}
	}

	// Get the user ID from the JWT token
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is missing
	if userID == "" {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user ID from JWT token",
		})
		return database.Team{}
	}

	// Retrieve the team from the database by name
	team, err := db.GetTeamByName(c, teamNameParam)

	// Check for errors when retrieving the team
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team from the name",
		})
		return database.Team{}
	}

	// Check if the user is the owner of the team
	if team.OwnerID != userID {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "You cannot modify this part of the team because you are not its owner",
		})
		return database.Team{}
	}

	return team
}

// VerifyOrgOwnership verifies if the user has ownership of an organization.
func VerifyOrgOwnership(c *gin.Context, token *jwt.Token, db *database.Queries) database.Org {
	// Get the organization name parameter
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Organization name is missing",
		})
		return database.Org{}
	}

	// Get the user ID from the JWT token
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is missing
	if userID == "" {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user ID from JWT token",
		})
		return database.Org{}
	}

	// Retrieve the organization from the database by name
	org, err := db.GetOrgByName(c, orgNameParam)

	// Check for errors when retrieving the organization
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get organization",
		})
		return database.Org{}
	}

	// Check if the user is the owner of the organization
	if org.OwnerID != userID {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "You cannot modify this part of the organization because you are not its owner",
		})
		return database.Org{}
	}

	return org
}

// VerifyOrgAndGetProjectByParamName verifies the organization ownership and retrieves a project by its name.
func VerifyOrgAndGetProjectByParamName(c *gin.Context) (database.Project, *database.Queries) {
	// Extract JWT token and database connection
	token, _, db := ExtractDBAndToken(c)

	// Check if the token is nil
	if token == nil {
		return database.Project{}, nil
	}

	// Verify organization ownership
	org := VerifyOrgOwnership(c, token, db)

	// Check if the organization ID is missing
	if org.ID == "" {
		return database.Project{}, nil
	}

	// Get projects by organization ID
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return database.Project{}, nil
	}

	// Get the project name parameter
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Project name is missing",
		})
		return database.Project{}, nil
	}

	// Find the project by name
	for _, project := range projects {
		if project.Name == projectNameParam {
			return project, db
		}
	}

	c.XML(http.StatusBadRequest, gin.H{
		"error": "Project name is missing",
	})
	return database.Project{}, nil
}

// GetProjectByParamName retrieves a project by its name and organization name.
func GetProjectByParamName(c *gin.Context, db *database.Queries) database.Project {
	// Get the organization name parameter
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Organization name is missing",
		})
		return database.Project{}
	}

	// Retrieve the organization by name
	org, err := db.GetOrgByName(c, orgNameParam)

	// Check for errors when retrieving the organization
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get organization",
		})
		return database.Project{}
	}

	// Get projects by organization ID
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return database.Project{}
	}

	// Get the project name parameter
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Project name is missing",
		})
		return database.Project{}
	}

	// Find the project by name
	for _, project := range projects {
		if project.Name == projectNameParam {
			return project
		}
	}

	return database.Project{}
}

// GetTaskByParamName retrieves a task by its name and validates ownership.
func GetTaskByParamName(c *gin.Context) (database.Task, *database.Queries) {
	// Get the team, database, and user ID
	team, db, userID := GetTeamAndDatabaseAndUserID(c)

	// Check if the team name is missing
	if team.Name == "" {
		return database.Task{}, nil
	}

	// Get the task name parameter
	taskNameParam := c.Param("subtaskName")

	// Check if the task name is missing
	if taskNameParam == "" {
		// Get the task name parameter
		taskNameParam = c.Param("taskName")

		// Check if the task name is missing
		if taskNameParam == "" {
			c.XML(http.StatusBadRequest, config.ErrorResponse{
				Message: "Task name is missing",
			})
			return database.Task{}, nil
		}
	}

	// Retrieve the task by name
	task, err := db.GetTaskByName(c, taskNameParam)

	// Check for errors when retrieving the task
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get task",
		})
		return database.Task{}, nil
	}

	var errors []string

	// Check if the user is the owner of the task
	if userID == task.OwnerID {
		return task, db
	} else {
		errors = append(errors, "User is not the task owner")
	}

	// Check if the user is the owner of the team
	if userID == team.OwnerID {
		return task, db
	} else {
		errors = append(errors, "User is not the team owner")
	}

	// Return unauthorized error if there are errors
	if len(errors) > 0 {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": errors[0],
		})
		return database.Task{}, nil
	}

	return task, db
}

// GetUserAndTeamByParam retrieves user, team, and database connections based on parameters in the request.
func GetUserAndTeamByParam(c *gin.Context) (database.User, database.Team, *database.Queries) {
	// Extract token, _, and db from the request context
	token, _, db := ExtractDBAndToken(c)

	// If no token is found, return empty User, Team, and nil Queries
	if token == nil {
		return database.User{}, database.Team{}, nil
	}

	// Define a struct to bind XML parameters to
	var params struct {
		User string `XML:"name"`
		Team string `XML:"name"`
	}

	// Bind XML parameters from the request to the params struct
	if err := c.ShouldBindXML(&params); err != nil {
		// Return a bad request response with an error message
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return database.User{}, database.Team{}, nil
	}

	// Verify team ownership based on parameters
	team := VerifyTeamOwnershipFromParam(c, token, db, params.Team)

	// If team ID is empty, return empty User, Team, and nil Queries
	if team.ID == "" {
		return database.User{}, database.Team{}, nil
	}

	// Retrieve the user by name from the database
	user, err := db.GetUserByName(c, params.User)
	if err != nil {
		// Return a bad request response with an error message
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return database.User{}, database.Team{}, nil
	}

	// Return the retrieved user, team, and database connection
	return user, team, db
}

// GetTeamAndProjectByParam retrieves project, team, and database connections based on parameters in the request.
func GetTeamAndProjectByParam(c *gin.Context) (database.Project, database.Team, *database.Queries) {
	// Extract token, _, and db from the request context
	token, _, db := ExtractDBAndToken(c)

	// If no token is found, return empty Project, Team, and nil Queries
	if token == nil {
		return database.Project{}, database.Team{}, nil
	}

	// Define a struct to bind XML parameters to
	var params struct {
		Project string `XML:"project"`
		Team    string `XML:"team"`
	}

	// Bind XML parameters from the request to the params struct
	if err := c.ShouldBindXML(&params); err != nil {
		// Return a bad request response with an error message
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return database.Project{}, database.Team{}, nil
	}

	// Retrieve the team by name from the database
	team, err := db.GetTeamByName(c, params.Team)

	// If there is an error or team ID is empty, return empty Project, Team, and nil Queries
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team from the name",
		})
		return database.Project{}, database.Team{}, nil
	}

	// Retrieve the project by name from the database
	project, err := db.GetProjectByName(c, params.Project)
	if err != nil {
		// Return a bad request response with an error message
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return database.Project{}, database.Team{}, nil
	}

	// Return the retrieved project, team, and database connection
	return project, team, db
}

// GetTeamAndDatabaseAndUserID retrieves team, database, and user ID based on the request and token.
func GetTeamAndDatabaseAndUserID(c *gin.Context) (database.Team, *database.Queries, string) {
	// Extract token, _, and db from the request context
	token, _, db := ExtractDBAndToken(c)

	// If no token is found, return empty Team, nil Queries, and an empty user ID
	if token == nil {
		return database.Team{}, nil, ""
	}

	// Extract user ID from the token claims
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// If user ID is empty, return an error response
	if userID == "" {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user ID from JWT token",
		})
		return database.Team{}, nil, ""
	}

	// Retrieve project based on parameter name
	project := GetProjectByParamName(c, db)

	// If project name is empty, return empty Team, nil Queries, and an empty user ID
	if project.Name == "" {
		return database.Team{}, nil, ""
	}

	// Retrieve team name parameter from the request
	teamNameParam := c.Param("teamName")

	// If team name parameter is empty, return an error response
	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Organization name is missing",
		})
		return database.Team{}, nil, ""
	}

	// Retrieve team by name from the database
	team, err := db.GetTeamByName(c, teamNameParam)

	// If there is an error, return an error response
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get organization",
		})
		return database.Team{}, nil, ""
	}

	// Retrieve all users from the team
	users, err := db.GetAllUsersFromTeam(c, team.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team users",
		})
		return database.Team{}, nil, ""
	}

	// Check if the user is in the team or the owner of the team
	userInTeam := false
	for _, user := range users {
		if userID == user.UserID {
			userInTeam = true
			break
		}
	}

	if userID == team.OwnerID {
		userInTeam = true
	}

	// If the user is not in the team, return an unauthorized response
	if !userInTeam {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "User is not in team",
		})
		return database.Team{}, nil, ""
	}

	// Return the retrieved team, database connection, and user ID
	return team, db, userID
}

// VerifyTeamMembershipAndGetProject verifies team membership and retrieves project, team, database, and user ID.
func VerifyTeamMembershipAndGetProject(c *gin.Context) (database.Project, database.Team, *database.Queries, string) {
	// Retrieve team, database, and user ID using the GetTeamAndDatabaseAndUserID function
	team, db, userID := GetTeamAndDatabaseAndUserID(c)

	// If team name is empty, return empty Project, Team, nil Queries, and an empty user ID
	if team.Name == "" {
		return database.Project{}, database.Team{}, nil, ""
	}

	// Retrieve project based on parameter name
	project := GetProjectByParamName(c, db)

	// If project name is empty, return empty Project, Team, nil Queries, and an empty user ID
	if project.Name == "" {
		return database.Project{}, database.Team{}, nil, ""
	}

	// Return the retrieved project, team, database connection, and user ID
	return project, team, db, userID
}

// SetProjectUpdatedTime updates the "UpdatedAt" field of a project in the database.
func SetProjectUpdatedTime(c *gin.Context, db *database.Queries, project database.Project) {
	// Define update arguments for the project
	args_project := database.UpdateProjectParams{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		UpdatedAt:   time.Now(),
	}

	// Update the project in the database
	_, err := db.UpdateProject(c, args_project)
	if err != nil {
		// Return an internal server error response with an error message
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to update project",
		})
		return
	}
}

// SetProjectUpdatedTimeUsingTask updates the "UpdatedAt" field of a project using a task.
func SetProjectUpdatedTimeUsingTask(c *gin.Context, db *database.Queries, task database.Task) database.Project {
	// Retrieve the project by ID from the database
	project, err := db.GetProjectByID(c, task.ProjectID)
	if err != nil {
		// Return an internal server error response with an error message
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get project",
		})
		return database.Project{}
	}

	// Call the SetProjectUpdatedTime function to update the project
	SetProjectUpdatedTime(c, db, project)

	// Return the updated project
	return project
}

// SetOrgUpdatedTime updates the "UpdatedAt" field of an organization in the database.
func SetOrgUpdatedTime(c *gin.Context, db *database.Queries, org database.Org) {
	// Define update arguments for the organization
	args_org := database.UpdateOrgParams{
		ID:          org.ID,
		Name:        org.Name,
		Description: org.Description,
		UpdatedAt:   time.Now(),
	}

	// Update the organization in the database
	_, err := db.UpdateOrg(c, args_org)
	if err != nil {
		// Return an internal server error response with an error message
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to update org",
		})
		return
	}
}

// SetOrgUpdatedTimeUsingProject updates the "UpdatedAt" field of an organization using a project.
func SetOrgUpdatedTimeUsingProject(c *gin.Context, db *database.Queries, project database.Project) {
	// Retrieve the organization by ID from the database
	org, err := db.GetOrgByID(c, project.OrgID)
	if err != nil {
		// Return an internal server error response with an error message
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get org",
		})
		return
	}

	// Call the SetOrgUpdatedTime function to update the organization
	SetOrgUpdatedTime(c, db, org)
}

// SetTeamUpdatedTime updates the "UpdatedAt" field of a team in the database.
func SetTeamUpdatedTime(c *gin.Context, db *database.Queries, teamID string) {
	// Retrieve the team by ID from the database
	team, err := db.GetTeamByID(c, teamID)
	if err != nil {
		// Return an internal server error response with an error message
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team from ID",
		})
		return
	}

	// Define update arguments for the team
	args_team := database.UpdateTeamParams{
		ID:          team.ID,
		Name:        team.Name,
		Description: team.Description,
		UpdatedAt:   time.Now(),
	}

	// Update the team in the database
	_, err = db.UpdateTeam(c, args_team)
	if err != nil {
		// Return an internal server error response with an error message
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to update team",
		})
		return
	}
}

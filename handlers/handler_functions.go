package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/tokenPackage"
)

// ExtractDBAndToken is a function that extracts a JWT token, its corresponding string,
// and a database connection from a Gin context.
func ExtractDBAndToken(c *gin.Context) (*jwt.Token, string, *database.Queries) {
	// Extract the JWT token string from the request header.
	tokenString := tokenPackage.ExtractJWTTokenFromHeader(c.Request)

	// Check if the token string is empty or missing.
	if tokenString == "" {
		// Return an unauthorized response and nil values.
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "JWT token is missing or invalid",
		})
		return nil, "", nil
	}

	// Parse and validate the JWT token.
	token, err := tokenPackage.ParseAndValidateJWTToken(tokenString)
	if err != nil {
		// Return an unauthorized response and nil values if the token is invalid.
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Invalid JWT token",
		})
		return nil, "", nil
	}

	// Check if there is a token with the same ID in the TokenMap and replace the token with it.
	for _, tok := range tokenPackage.TokenMap {
		if tok.Claims.(jwt.MapClaims)["ID"] == token.Claims.(jwt.MapClaims)["ID"] {
			token = tok
			break
		}
	}

	// Check if the token has been revoked.
	if token.Claims.(jwt.MapClaims)["Revoked"] == true {
		// Return an unauthorized response and nil values if the token is revoked.
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Trying to use revoked JWT token",
		})
		return nil, "", nil
	}

	// Retrieve the database connection from the Gin context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// Return an internal server error response and nil values if the database connection is unavailable.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return nil, "", nil
	}

	// Return the validated token, token string, and database connection.
	return token, tokenString, db
}

func VerifyTeamOwnershipFromParam(c *gin.Context, token *jwt.Token, db *database.Queries, teamNameParam string) database.Team {
	// Check if the team name is missing. If so, return a Bad Request response.
	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Team name is missing",
		})
		return database.Team{}
	}

	// Extract the user ID from the JWT token claims.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is empty. If so, return an Internal Server Error response.
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return database.Team{}
	}

	// Retrieve a team with a certain name from the database.
	team, err := db.GetTeamByName(c, teamNameParam)
	// Check for errors while fetching the team from the database.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get team",
		})
		return database.Team{}
	}

	// Check if the current user is the owner of the team. If not, return an Unauthorized response.
	if team.OwnerID != userID {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "You cannot modify this part of the team because you are not its owner",
		})
		return database.Team{}
	}

	// Return the team if the user has ownership rights.
	return team
}

func VerifyOrgOwnership(c *gin.Context, token *jwt.Token, db *database.Queries) database.Org {
	// Get the organization name from the parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing. If so, return a Bad Request response.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return database.Org{}
	}

	// Extract the user ID from the JWT token claims.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is empty. If so, return an Internal Server Error response.
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return database.Org{}
	}

	// Retrieve an organization with a certain name from the database.
	org, err := db.GetOrgByName(c, orgNameParam)
	// Check for errors while fetching the organization from the database.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return database.Org{}
	}

	// Check if the current user is the owner of the organization. If not, return an Unauthorized response.
	if org.OwnerID != userID {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "You cannot modify this part of the organization because you are not its owner",
		})
		return database.Org{}
	}

	// Return the organization if the user has ownership rights.
	return org
}

func VerifyOrgAndGetProjectByParamName(c *gin.Context) (database.Project, *database.Queries) {
	// Extract the authentication token, database connection, and user from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the authentication token is not present, return early.
	if token == nil {
		return database.Project{}, nil
	}

	// Verify that the user owns the organization based on the token.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization is not found or the user doesn't own it, return early.
	if org.ID == "" {
		return database.Project{}, nil
	}

	// Retrieve a list of projects from the database.
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		// If there was an error getting the projects from the database, return a 500 Internal Server Error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return database.Project{}, nil
	}

	// Get the project name from the parameter.
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing. If so, return a Bad Request response.
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Project name is missing",
		})
		return database.Project{}, nil
	}

	// Loop through the projects and check if the project name matches the name in the parameter.
	for _, project := range projects {
		if project.Name == projectNameParam {
			// If the project name matches, return the project and the database connection.
			return project, db
		}
	}

	return database.Project{}, nil
}

func GetProjectByParamName(c *gin.Context, db *database.Queries) database.Project {
	// Get the organization name from the parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing. If so, return a Bad Request response.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return database.Project{}
	}

	// Retrieve an organization with a certain name from the database.
	org, err := db.GetOrgByName(c, orgNameParam)
	// Check for errors while fetching the organization from the database.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return database.Project{}
	}

	// Retrieve a list of projects from the database.
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		// If there was an error getting the projects from the database, return a 500 Internal Server Error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return database.Project{}
	}

	// Get the project name from the parameter.
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing. If so, return a Bad Request response.
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Project name is missing",
		})
		return database.Project{}
	}

	// Loop through the projects and check if the project name matches the name in the parameter.
	for _, project := range projects {
		if project.Name == projectNameParam {
			// If the project name matches, return the project and the database connection.
			return project
		}
	}

	return database.Project{}
}

func GetTaskByParamName(c *gin.Context) (database.Task, *database.Queries) {
	// Extract the team, database connection, and user ID from the context.
	team, db, userID := GetTeamAndDatabaseAndUserID(c)

	// If the team is not found or the user doesn't own it, return early.
	if team.Name == "" {
		return database.Task{}, nil
	}

	taskNameParam := c.Param("taskName")

	// Check if the project name is missing. If so, return a Bad Request response.
	if taskNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Task name is missing",
		})
		return database.Task{}, nil
	}

	// Retrieve the task from the database.
	task, err := db.GetTaskByName(c, taskNameParam)
	if err != nil {
		// If there is an error retrieving the task, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get task",
		})
		return database.Task{}, nil
	}

	var errors []string

	// Check if the current user is the owner of the task. If not, return an Unauthorized response.
	if userID == task.OwnerID {
		return task, db
	} else {
		errors = append(errors, "User is not the task owner")
	}

	// Check if the current user is the owner of the task. If not, return an Unauthorized response.
	if userID == team.OwnerID {
		return task, db
	} else {
		errors = append(errors, "User is not the team owner")
	}

	if len(errors) > 0 {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": errors[0],
		})
		return database.Task{}, nil
	}

	return task, db
}

func GetUserAndTeamByParam(c *gin.Context) (database.User, database.Team, *database.Queries) {
	// Extract the authentication token, database connection, and user from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the authentication token is not present, return early.
	if token == nil {
		return database.User{}, database.Team{}, nil
	}

	// Define a struct to hold the parameters received in the XML request.
	var params struct {
		User string `XML:"name"`
		Team string `XML:"name"`
	}

	// Bind the XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		// If there is an error binding the XML, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return database.User{}, database.Team{}, nil
	}

	// Verify that the user owns the team based on the token.
	team := VerifyTeamOwnershipFromParam(c, token, db, params.Team)

	// If the team is not found or the user doesn't own it, return early.
	if team.ID == "" {
		return database.User{}, database.Team{}, nil
	}

	user, err := db.GetUserByName(c, params.User)
	if err != nil {
		// If there is an error binding the XML, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return database.User{}, database.Team{}, nil
	}

	return user, team, db
}

func GetTeamAndProjectByParam(c *gin.Context) (database.Project, database.Team, *database.Queries) {
	// Extract the authentication token, database connection, and user from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the authentication token is not present, return early.
	if token == nil {
		return database.Project{}, database.Team{}, nil
	}

	// Define a struct to hold the parameters received in the XML request.
	var params struct {
		Project string `XML:"name"`
		Team    string `XML:"name"`
	}

	// Bind the XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		// If there is an error binding the XML, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return database.Project{}, database.Team{}, nil
	}

	// Retrieve a team with a certain name from the database.
	team, err := db.GetTeamByName(c, params.Team)
	// Check for errors while fetching the team from the database.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get team",
		})
		return database.Project{}, database.Team{}, nil
	}

	// If the team is not found or the user doesn't own it, return early.
	if team.ID == "" {
		return database.Project{}, database.Team{}, nil
	}

	project, err := db.GetProjectByName(c, params.Project)
	if err != nil {
		// If there is an error binding the XML, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return database.Project{}, database.Team{}, nil
	}

	return project, team, db
}

func GetTeamAndDatabaseAndUserID(c *gin.Context) (database.Team, *database.Queries, string) {
	// Extract the authentication token, database connection, and user from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the authentication token is not present, return early.
	if token == nil {
		return database.Team{}, nil, ""
	}

	// Extract the user ID from the JWT token claims.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is empty. If so, return an Internal Server Error response.
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return database.Team{}, nil, ""
	}

	// Get the project from the context.
	project := GetProjectByParamName(c, db)

	// If project has no name, return early without further processing.
	if project.Name == "" {
		return database.Team{}, nil, ""
	}

	// Get the organization name from the parameter.
	teamNameParam := c.Param("teamName")

	// Check if the organization name is missing. If so, return a Bad Request response.
	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return database.Team{}, nil, ""
	}

	// Retrieve an organization with a certain name from the database.
	team, err := db.GetTeamByName(c, teamNameParam)
	// Check for errors while fetching the organization from the database.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return database.Team{}, nil, ""
	}

	// Get all the users in the team
	users, err := db.GetAllUsersFromTeam(c, team.ID)
	if err != nil {
		// If there is an error while getting the team, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get team users",
		})
		return database.Team{}, nil, ""
	}

	userInTeam := false
	// Check if the user is in the team
	for _, user := range users {
		if user.UserID == userID || team.OwnerID == userID {
			userInTeam = true
			break
		}
	}
	if !userInTeam {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "User is not in team",
		})
		return database.Team{}, nil, ""
	}

	return team, db, userID
}

func VerifyTeamMembershipAndGetProject(c *gin.Context) (database.Project, database.Team, *database.Queries, string) {
	// Extract the team and database connection from the context.
	team, db, userID := GetTeamAndDatabaseAndUserID(c)

	// If the team is not found or the user doesn't own it, return early.
	if team.Name == "" {
		return database.Project{}, database.Team{}, nil, ""
	}

	// Get the project from the context.
	project := GetProjectByParamName(c, db)

	// If project has no name, return early without further processing.
	if project.Name == "" {
		return database.Project{}, database.Team{}, nil, ""
	}

	return project, team, db, userID
}

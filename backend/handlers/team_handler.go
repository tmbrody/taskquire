package handlers

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/tokenPackage"
)

// CreateTeamHandler handles the creation of a new team.
func CreateTeamHandler(c *gin.Context) {
	// Extract the JWT token, database connection, and user info.
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	// Define parameters for the new team.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Parse XML request body into the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Get the user ID from the JWT token.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)
	if userID == "" {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user ID from JWT token",
		})
		return
	}

	// Generate a new team ID.
	teamID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Unable to generate team ID",
		})
		return
	}

	// Prepare team description.
	teamDescriptionString := params.Description
	var teamDescription sql.NullString
	if teamDescriptionString != "" {
		teamDescription.String = teamDescriptionString
		teamDescription.Valid = true
	} else {
		teamDescription.String = params.Description
		teamDescription.Valid = false
	}

	// Create arguments for creating the team.
	argsCreate := database.CreateTeamParams{
		ID:          teamID.String(),
		Name:        params.Name,
		Description: teamDescription,
		OwnerID:     userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Call the database function to create the team.
	_, err = db.CreateTeam(c, argsCreate)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to create new team",
		})
		return
	}

	// Create arguments for adding the user to the team.
	args := database.AddUsertoTeamParams{
		UserID: userID,
		TeamID: teamID.String(),
	}

	// Call the database function to add the user to the team.
	_, err = db.AddUsertoTeam(c, args)
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to add creator to %s", params.Name),
		})
		return
	}

	// Return the created team as XML response.
	c.XML(http.StatusCreated, argsCreate)
}

// GetTeamsHandler handles retrieving a list of all teams.
func GetTeamsHandler(c *gin.Context) {
	// Retrieve the database connection.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	var token *jwt.Token

	// Extract JWT token from the request header
	tokenString := tokenPackage.ExtractJWTTokenFromHeader(c.Request)

	// Check if the token is missing or invalid
	if tokenString != "" {
		token, _, _ = ExtractDBAndToken(c)
	}

	var teams []database.Team
	var allTeams []database.TeamsUser
	var err error

	// If the token is nil, return.
	if token == nil {
		// Get all organizations from the database.
		teams, err = db.GetAllTeams(c)

		// If there's an error retrieving organizations, return an error.
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get teams",
			})
			return
		}
	} else {
		// Get all organizations from the database.
		allTeams, err = db.GetAllTeamsByUser(c, token.Claims.(jwt.MapClaims)["Subject"].(string))

		// If there's an error retrieving organizations, return an error.
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get teams",
			})
			return
		}

		// Prepare a slice to store team information.
		for _, team := range allTeams {
			t, err := db.GetTeamByID(c, team.TeamID)
			if err != nil {
				c.XML(http.StatusInternalServerError, config.ErrorResponse{
					Message: "Unable to get team from ID",
				})
				return
			}
			// Add team information to the slice.
			teams = append(teams, t)
		}
	}

	userList := []string{}
	projectList := []string{}

	// Prepare a slice to store team information.
	var teamMap []gin.H
	for _, team := range teams {
		// Retrieve users and projects associated with the team.
		users, _ := db.GetAllUsersFromTeam(c, team.ID)
		projects, _ := db.GetAllProjectsFromTeam(c, team.ID)

		// Prepare a list of user names in the team.
		for _, user := range users {
			u, _ := db.GetUserByID(c, user.UserID)
			userList = append(userList, u.Name)
		}

		// Prepare a list of project names in the team.
		for _, project := range projects {
			p, _ := db.GetProjectByID(c, project.ProjectID)
			projectList = append(projectList, p.Name)
		}

		// Retrieve the team owner.
		owner, _ := db.GetUserByID(c, team.OwnerID)

		// Add team information to the slice.
		teamMap = append(teamMap, gin.H{
			"ID":          team.ID,
			"Name":        team.Name,
			"Description": team.Description,
			"OwnerName":   owner.Name,
			"CreatedAt":   team.CreatedAt,
			"UpdatedAt":   team.UpdatedAt,
			"Users": func() string {
				if len(userList) == 0 {
					return "None"
				}
				return strings.Join(userList, ", ")
			}(),
			"Projects": func() string {
				if len(projectList) == 0 {
					return "None"
				}
				return strings.Join(projectList, ", ")
			}(),
		})
	}

	xmlData, err := xml.Marshal(gin.H{"Teams": teamMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "team")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetOneTeamHandler handles retrieving information about a single team.
func GetOneTeamHandler(c *gin.Context) {
	// Retrieve the database connection.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Retrieve the team name from the URL parameter.
	teamNameParam := c.Param("teamName")

	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Team name is missing",
		})
		return
	}

	// Retrieve team information from the database.
	team, err := db.GetTeamByName(c, teamNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team from name",
		})
		return
	}

	// Retrieve users associated with the team.
	users, err := db.GetAllUsersFromTeam(c, team.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team users",
		})
		return
	}

	// Prepare a list of user names in the team.
	var userNames []string
	for _, user := range users {
		u, err := db.GetUserByID(c, user.UserID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get team user name",
			})
			return
		}
		userNames = append(userNames, u.Name)
	}

	if userNames == nil {
		userNames = append(userNames, "No users in team")
	}

	// Create a UserList struct to store user names.
	userList := UserList{Users: userNames}

	// Retrieve the team owner.
	owner, err := db.GetUserByID(c, team.OwnerID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team owner name",
		})
		return
	}

	// Prepare team information for response.
	var teamMap []gin.H
	teamMap = append(teamMap, gin.H{
		"ID":          team.ID,
		"Name":        team.Name,
		"Description": team.Description,
		"OwnerName":   owner.Name,
		"CreatedAt":   team.CreatedAt,
		"UpdatedAt":   team.UpdatedAt,
		"Users":       userList,
	})

	xmlData, err := xml.Marshal(gin.H{"Teams": teamMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "team")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetProjectsFromTeamHandler handles retrieving a list of projects associated with a team.
func GetProjectsFromTeamHandler(c *gin.Context) {
	// Get the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the team name from the URL parameter.
	teamNameParam := c.Param("teamName")

	// Check if the team name is missing.
	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Team name is missing",
		})
		return
	}

	// Get the team information from the database.
	team, err := db.GetTeamByName(c, teamNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team",
		})
		return
	}

	// Get the projects associated with the team.
	projectIDs, err := db.GetAllProjectsFromTeam(c, team.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return
	}

	var projects []database.Project

	// Prepare a slice to store project information.
	for _, project := range projectIDs {
		p, err := db.GetProjectByID(c, project.ProjectID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get project from ID",
			})
			return
		}

		// Add the project information to the slice.
		projects = append(projects, p)
	}

	teams := []string{}

	// Create a list of project information to be returned.
	var projectMap []gin.H
	for _, project := range projects {
		// Get the organization associated with the project.
		org, _ := db.GetOrgByID(c, project.OrgID)

		// Get the teams associated with the project.
		teamIDs, _ := db.GetAllTeamsFromProject(c, project.ID)
		for _, team := range teamIDs {
			t, _ := db.GetTeamByID(c, team.TeamID)
			teams = append(teams, t.Name)
		}

		projectMap = append(projectMap, gin.H{
			"ID":          project.ID,
			"Name":        project.Name,
			"Description": project.Description,
			"OrgName":     org.Name,
			"CreatedAt":   project.CreatedAt,
			"UpdatedAt":   project.UpdatedAt,
			"Teams": func() string {
				if len(teams) == 0 {
					return "None"
				}
				return strings.Join(teams, ", ")
			}(),
		})
	}

	xmlData, err := xml.Marshal(gin.H{"Projects": projectMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "project")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetTeamsFromProjectHandler handles retrieving a list of teams associated with a project.
func GetTeamsFromProjectHandler(c *gin.Context) {
	// Get the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the project name from the URL parameter.
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing.
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Project name is missing",
		})
		return
	}

	// Get the project information from the database.
	project, err := db.GetProjectByName(c, projectNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get project",
		})
		return
	}

	// Get the teams associated with the project.
	teamIDs, err := db.GetAllTeamsFromProject(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return
	}

	var teams []database.Team

	for _, team := range teamIDs {
		// Get the team information from the database.
		t, err := db.GetTeamByID(c, team.TeamID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get team from ID",
			})
			return
		}

		// Add the team information to the slice.
		teams = append(teams, t)
	}

	// Create a list of team information to be returned.
	var teamMap []gin.H
	for _, team := range teams {
		// Get the owner of the team
		owner, _ := db.GetUserByID(c, team.OwnerID)

		taskList := []string{}

		// Retrieve all tasks for the project from the database
		taskProjectIDs, _ := db.GetTasksByProjectID(c, project.ID)

		// Retrieve all tasks for the team from the database
		taskTeamIDs, _ := db.GetTasksByTeamID(c, team.ID)

		// If the task is in both lists, add it to the task list
		for _, pTask := range taskProjectIDs {
			for _, tTask := range taskTeamIDs {
				if pTask.ID == tTask.ID {
					task, _ := db.GetTaskByID(c, pTask.ID)
					if task.ParentID == "" {
						taskList = append(taskList, task.Name)
					}
				}
			}
		}

		teamMap = append(teamMap, gin.H{
			"ID":          team.ID,
			"Name":        team.Name,
			"Description": team.Description,
			"OwnerName":   owner.Name,
			"CreatedAt":   team.CreatedAt,
			"UpdatedAt":   team.UpdatedAt,
			"Tasks": func() string {
				if len(taskList) == 0 {
					return "None"
				}
				return strings.Join(taskList, ", ")
			}(),
		})
	}

	xmlData, err := xml.Marshal(gin.H{"Teams": teamMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "team")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// UpdateTeamHandler handles updating team information.
func UpdateTeamHandler(c *gin.Context) {
	// Extract the JWT token, database connection, and user info.
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	// Retrieve the team name from the URL parameter.
	teamNameParam := c.Param("teamName")

	// Verify team ownership based on the token and team name.
	team := VerifyTeamOwnershipFromParam(c, token, db, teamNameParam)

	if team.ID == "" {
		return
	}

	// Define parameters for updating the team.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Parse XML request body into the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Determine the updated team name.
	teamName := params.Name
	if teamName == "" {
		teamName = team.Name
	}

	// Prepare team description.
	teamDescriptionString := params.Description
	var teamDescription sql.NullString
	if teamDescriptionString != "" {
		teamDescription.String = teamDescriptionString
		teamDescription.Valid = true
	} else {
		teamDescription.String = params.Description
		teamDescription.Valid = false
	}

	// Create arguments for updating the team.
	args := database.UpdateTeamParams{
		ID:          team.ID,
		Name:        teamName,
		Description: teamDescription,
		UpdatedAt:   time.Now(),
	}

	// Call the database function to update the team.
	_, err := db.UpdateTeam(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to update %s", team.Name),
		})
		return
	}

	// Return the updated team information as an XML response.
	c.XML(http.StatusOK, args)
}

// AddUserToTeamHandler handles adding a user to a team.
func AddUserToTeamHandler(c *gin.Context) {
	// Retrieve the user, team, and database from the URL parameters.
	user, team, db := GetUserAndTeamByParam(c)

	if user.ID == "" {
		return
	}

	// Create arguments for adding the user to the team.
	args := database.AddUsertoTeamParams{
		UserID: user.ID,
		TeamID: team.ID,
	}

	// Call the database function to add the user to the team.
	_, err := db.AddUsertoTeam(c, args)
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to add %s to %s", user.Name, team.Name),
		})
		return
	}

	// Return a success message as an XML response.
	c.XML(http.StatusOK, fmt.Sprintf("%s has been added to %s", user.Name, team.Name))
}

// RemoveUserFromTeamHandler handles removing a user from a team.
func RemoveUserFromTeamHandler(c *gin.Context) {
	// Retrieve the user, team, and database from the URL parameters.
	user, team, db := GetUserAndTeamByParam(c)

	if user.ID == "" {
		return
	}

	// Prepare arguments for checking if the user is in the team.
	argsGet := database.GetOneUserFromTeamParams{
		UserID: user.ID,
		TeamID: team.ID,
	}

	// Check if the user is in the team.
	_, err := db.GetOneUserFromTeam(c, argsGet)
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to find %s in %s", user.Name, team.Name),
		})
		return
	}

	// Prepare arguments for removing the user from the team.
	argsRemove := database.RemoveUserFromTeamParams{
		UserID: user.ID,
		TeamID: team.ID,
	}

	// Call the database function to remove the user from the team.
	_, err = db.RemoveUserFromTeam(c, argsRemove)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to remove %s from %s", user.Name, team.Name),
		})
		return
	}

	// Return a success message as an XML response.
	c.XML(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s has been removed from %s", user.Name, team.Name),
	})
}

// DeleteTeamHandler handles deleting a team.
func DeleteTeamHandler(c *gin.Context) {
	// Extract the JWT token, database connection, and user info.
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	// Retrieve the team name from the URL parameter.
	teamNameParam := c.Param("teamName")

	// Verify team ownership based on the token and team name.
	team := VerifyTeamOwnershipFromParam(c, token, db, teamNameParam)

	if team.ID == "" {
		return
	}

	// Call the database function to delete the team.
	_, err := db.DeleteTeam(c, team.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to delete %s", team.Name),
		})
		return
	}

	// Return a success message as an XML response.
	c.XML(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s has been deleted", team.Name),
	})
}

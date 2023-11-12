package handlers

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"net/http"
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
		teams, err = db.GetTeamsByOwnerID(c, token.Claims.(jwt.MapClaims)["Subject"].(string))

		// If there's an error retrieving organizations, return an error.
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get teams",
			})
			return
		}
	}

	// Prepare a slice to store team information.
	var teamMap []gin.H
	for _, team := range teams {
		// Add team information to the slice.
		teamMap = append(teamMap, gin.H{
			"ID":          team.ID,
			"Name":        team.Name,
			"Description": team.Description,
			"OwnerID":     team.OwnerID,
			"CreatedAt":   team.CreatedAt,
			"UpdatedAt":   team.UpdatedAt,
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

	// Get the projects associated with the team.
	projectIDs, err := db.GetAllProjectsFromTeam(c, team.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return
	}

	// Prepare a list of user names in the team.
	var projectNames []string
	for _, project := range projectIDs {
		p, err := db.GetProjectByID(c, project.ProjectID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get project name",
			})
			return
		}
		projectNames = append(projectNames, p.Name)
	}

	// Create a UserList struct to store user names.
	projectList := ProjectList{Projects: projectNames}

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

	// Create a UserList struct to store user names.
	userList := UserList{Users: userNames}

	// Prepare team information for response.
	var teamMap []gin.H
	teamMap = append(teamMap, gin.H{
		"ID":          team.ID,
		"Name":        team.Name,
		"Description": team.Description,
		"OwnerID":     team.OwnerID,
		"CreatedAt":   team.CreatedAt,
		"UpdatedAt":   team.UpdatedAt,
		"Projects":    projectList,
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
			Message: "Unable to get organization",
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
		teamMap = append(teamMap, gin.H{
			"ID":          team.ID,
			"Name":        team.Name,
			"Description": team.Description,
			"OwnerID":     team.OwnerID,
			"CreatedAt":   team.CreatedAt,
			"UpdatedAt":   team.UpdatedAt,
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

package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

// CreateTeamHandler handles the creation of a new team based on XML input data.
func CreateTeamHandler(c *gin.Context) {
	// Extract token, database connection, and user information from the request context.
	token, _, db := ExtractDBAndToken(c)

	// Check if the token is missing.
	if token == nil {
		return
	}

	// Define a struct to hold the XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML data from the request to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		// If there's an error while binding XML data, return a bad request error.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Extract the user ID from the JWT token's claims.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if user ID couldn't be extracted from the token.
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

	// Generate a new UUID for the team ID.
	teamID, err := uuid.NewUUID()
	if err != nil {
		// If there's an error while generating a UUID, return a bad request error.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate team ID",
		})
		return
	}

	// Get the task description from the params or use the current task description if it's empty.
	teamDescriptionString := params.Description

	var teamDescription sql.NullString

	if teamDescriptionString != "" {
		teamDescription.String = teamDescriptionString
		teamDescription.Valid = true
	} else {
		teamDescription.String = params.Description
		teamDescription.Valid = false
	}

	// Create arguments for creating a new team in the database.
	args := database.CreateTeamParams{
		ID:          teamID.String(),
		Name:        params.Name,
		Description: teamDescription,
		OwnerID:     userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Attempt to create the new team in the database.
	_, err = db.CreateTeam(c, args)
	if err != nil {
		// If there's an error while creating the team, return an internal server error.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new team",
		})
		return
	}

	// Return a success response with the created team data.
	c.XML(http.StatusCreated, args)
}

// GetTeamsHandler is a handler function that retrieves a list of teams from a database and responds with XML data.
func GetTeamsHandler(c *gin.Context) {
	// Attempt to retrieve the database connection from the Gin context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, respond with an internal server error message in XML format
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Retrieve a list of teams from the database
	teams, err := db.GetAllTeams(c)
	if err != nil {
		// If there's an error while getting teams from the database, respond with an internal server error message in XML format
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get teams",
		})
		return
	}

	// Create a slice to hold the team data in a map format
	var teamMap []gin.H
	for _, team := range teams {
		// Append team information to the teamMap slice
		teamMap = append(teamMap, gin.H{
			"ID":          team.ID,
			"Name":        team.Name,
			"Description": team.Description,
			"OwnerID":     team.OwnerID,
			"CreatedAt":   team.CreatedAt,
			"UpdatedAt":   team.UpdatedAt,
		})
	}

	// Respond with the list of teams in XML format with a 200 OK status
	c.XML(http.StatusOK, gin.H{
		"Teams": teamMap,
	})
}

// GetOneTeamHandler is a handler function that retrieves a single team based on a name from the database and responds with XML data.
func GetOneTeamHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Get the team name from the parameter.
	teamNameParam := c.Param("teamName")

	// Check if the team name is missing. If so, return a Bad Request response.
	if teamNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Team name is missing",
		})
		return
	}

	// Retrieve a team with a certain name from the database
	team, err := db.GetTeamByName(c, teamNameParam)
	if err != nil {
		// If there is an error while getting the team, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get team",
		})
		return
	}

	// Prepare the result data to be sent back as an XML response
	result := database.Team{
		Name:        team.Name,
		Description: team.Description,
		OwnerID:     team.OwnerID,
		CreatedAt:   team.CreatedAt,
		UpdatedAt:   team.UpdatedAt,
	}

	// Return the public information of the team as an XML response
	c.XML(http.StatusOK, result)
}

// UpdateTeamHandler handles requests to update team information.
func UpdateTeamHandler(c *gin.Context) {
	// Extract the authentication token, database connection, and user from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the authentication token is not present, return early.
	if token == nil {
		return
	}

	// Verify that the user owns the team based on the token.
	team := VerifyTeamOwnership(c, token, db)

	// If the team is not found or the user doesn't own it, return early.
	if team.ID == "" {
		return
	}

	// Define a struct to hold the parameters received in the XML request.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind the XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		// If there is an error binding the XML, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Get the team name from the params or use the current team name if it's empty.
	teamName := params.Name
	if teamName == "" {
		teamName = team.Name
	}

	// Get the task description from the params or use the current task description if it's empty.
	teamDescriptionString := params.Description

	var teamDescription sql.NullString

	if teamDescriptionString != "" {
		teamDescription.String = teamDescriptionString
		teamDescription.Valid = true
	} else {
		teamDescription.String = params.Description
		teamDescription.Valid = false
	}

	// Prepare the arguments for updating the team in the database.
	args := database.UpdateTeamParams{
		ID:          team.ID,
		Name:        teamName,
		Description: teamDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the team in the database.
	_, err := db.UpdateTeam(c, args)
	if err != nil {
		// If there is an error updating the team, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to update team",
		})
		return
	}

	// Respond with a success status and the updated team information.
	c.XML(http.StatusOK, args)
}

// DeleteTeamHandler is a function that handles the deletion of an team.
func DeleteTeamHandler(c *gin.Context) {
	// Extract the token, user ID, and database connection from the request context.
	token, _, db := ExtractDBAndToken(c)

	// If there's no token, return early without further processing.
	if token == nil {
		return
	}

	// Verify ownership of the team associated with the token.
	team := VerifyTeamOwnership(c, token, db)

	// If the team does not exist or there's an issue with ownership verification, return early.
	if team.ID == "" {
		return
	}

	// Attempt to delete the team from the database.
	_, err := db.DeleteTeam(c, team.ID)
	// If there's an error during the deletion process, return an error response.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete team",
		})
		return
	}

	// Respond with a success status and a message indicating the team was deleted.
	c.XML(http.StatusOK, "Team has been deleted successfully")
}

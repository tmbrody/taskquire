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

func VerifyTeamOwnership(c *gin.Context, token *jwt.Token, db *database.Queries) database.Team {
	// Get the team name from the parameter.
	teamNameParam := c.Param("teamName")

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

func GetProjectByParamName(c *gin.Context) (database.Project, *database.Queries) {
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

	// Get the project name from the parameter.
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing. If so, return a Bad Request response.
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Project name is missing",
		})
		return database.Project{}, nil
	}

	// Use the project name to get the project from the database.
	project, err := db.GetProjectByName(c, projectNameParam)
	if err != nil {
		c.XML(http.StatusNotFound, gin.H{
			"error": "Unable to get project",
		})
		return database.Project{}, nil
	}

	return project, db
}

func GetTaskByParamName(c *gin.Context) (database.Task, *database.Queries) {
	// Get the project from the context.
	project, db := GetProjectByParamName(c)

	// If project has no name, return early without further processing.
	if project.Name == "" {
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

	return task, db
}

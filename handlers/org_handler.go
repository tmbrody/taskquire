package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateOrgHandler(c *gin.Context) {
	// Extract token, database connection, and user information from the request context.
	token, _, db := ExtractDBAndToken(c)

	// Check if the token is missing.
	if token == nil {
		return
	}

	// Define a struct to hold parameters extracted from XML request.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML request body to the 'params' struct.
	if err := c.ShouldBindXML(&params); err != nil {
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

	// Generate a new UUID as the organization ID.
	orgID, err := uuid.NewUUID()

	// Check if there was an error generating the organization ID.
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate organization ID",
		})
		return
	}

	// Prepare arguments for creating a new organization in the database.
	args := database.CreateOrgParams{
		ID:          orgID.String(),
		Name:        params.Name,
		Description: params.Description,
		OwnerID:     userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Create a new organization in the database.
	_, err = db.CreateOrg(c, args)

	// Check if there was an error creating the organization.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new organization",
		})
		return
	}

	// Respond with the created organization details.
	c.XML(http.StatusCreated, args)
}

func GetOrgsHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Retrieve all organizations from the database
	orgs, err := db.GetAllOrgs(c)
	if err != nil {
		// If there is an error while getting organizations, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organizations",
		})
		return
	}

	// Map the organization data into a slice of maps for XML serialization
	var orgMap []gin.H
	for _, org := range orgs {
		orgMap = append(orgMap, gin.H{
			"ID":          org.ID,
			"Name":        org.Name,
			"Description": org.Description,
			"OwnerID":     org.OwnerID,
			"CreatedAt":   org.CreatedAt,
			"UpdatedAt":   org.UpdatedAt,
		})
	}

	// Return the list of organizations as an XML response
	c.XML(http.StatusOK, gin.H{
		"Orgs": orgMap,
	})
}

// GetOneOrgHandler is a handler function that retrieves a single organization based on a name from the database and responds with XML data.
func GetOneOrgHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Get the organization name from the parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing. If so, return a Bad Request response.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return
	}

	// Retrieve an organization with a certain name from the database
	org, err := db.GetOrgByName(c, orgNameParam)
	if err != nil {
		// If there is an error while getting the organization, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return
	}

	// Prepare the result data to be sent back as an XML response
	result := database.Org{
		Name:        org.Name,
		Description: org.Description,
		OwnerID:     org.OwnerID,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	}

	// Return the public information of the organization as an XML response
	c.XML(http.StatusOK, result)
}

// UpdateOrgHandler handles requests to update organization information.
func UpdateOrgHandler(c *gin.Context) {
	// Extract the authentication token, database connection, and user from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the authentication token is not present, return early.
	if token == nil {
		return
	}

	// Verify that the user owns the organization based on the token.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization is not found or the user doesn't own it, return early.
	if org.ID == "" {
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

	// Get the organization name from the params or use the current organization name if it's empty.
	orgName := params.Name
	if orgName == "" {
		orgName = org.Name
	}

	// Get the organization description from the params or use the current organization description if it's empty.
	orgDescription := params.Description
	if orgDescription == "" {
		orgDescription = org.Description
	}

	// Prepare the arguments for updating the organization in the database.
	args := database.UpdateOrgParams{
		ID:          org.ID,
		Name:        orgName,
		Description: orgDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the organization in the database.
	_, err := db.UpdateOrg(c, args)
	if err != nil {
		// If there is an error updating the organization, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to update organization",
		})
		return
	}

	// Respond with a success status and the updated organization information.
	c.XML(http.StatusOK, args)
}

// DeleteOrgHandler is a function that handles the deletion of an organization.
func DeleteOrgHandler(c *gin.Context) {
	// Extract the token, user ID, and database connection from the request context.
	token, _, db := ExtractDBAndToken(c)

	// If there's no token, return early without further processing.
	if token == nil {
		return
	}

	// Verify ownership of the organization associated with the token.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization does not exist or there's an issue with ownership verification, return early.
	if org.ID == "" {
		return
	}

	// Attempt to delete the organization from the database.
	_, err := db.DeleteOrg(c, org.ID)
	// If there's an error during the deletion process, return an error response.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete organization",
		})
		return
	}

	// Respond with a success status and a message indicating the project was deleted.
	c.XML(http.StatusOK, "Organization has been deleted successfully")
}

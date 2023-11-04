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

// CreateOrgHandler handles the creation of a new organization.
func CreateOrgHandler(c *gin.Context) {
	// Extract token, database connection, and user ID from the context.
	token, _, db := ExtractDBAndToken(c)

	// If the token is nil, return.
	if token == nil {
		return
	}

	// Define a struct to hold XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML request to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Get the user ID from the JWT token.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// If the user ID is empty, return an error.
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

	// Generate a new organization ID.
	orgID, err := uuid.NewUUID()

	// If there's an error generating the ID, return an error.
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate organization ID",
		})
		return
	}

	// Create arguments for creating a new organization.
	args := database.CreateOrgParams{
		ID:          orgID.String(),
		Name:        params.Name,
		Description: params.Description,
		OwnerID:     userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Create the new organization in the database.
	_, err = db.CreateOrg(c, args)

	// If there's an error creating the organization, return an error.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create a new organization",
		})
		return
	}

	// Return the created organization data.
	c.XML(http.StatusCreated, args)
}

// GetOrgsHandler retrieves a list of all organizations.
func GetOrgsHandler(c *gin.Context) {
	// Retrieve the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)

	// If there's an error retrieving the database connection, return an error.
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get a database connection",
		})
		return
	}

	// Get all organizations from the database.
	orgs, err := db.GetAllOrgs(c)

	// If there's an error retrieving organizations, return an error.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organizations",
		})
		return
	}

	// Create a list of organization data to return.
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

	// Return the list of organizations.
	c.XML(http.StatusOK, gin.H{
		"Orgs": orgMap,
	})
}

// GetOneOrgHandler retrieves information about a specific organization by its name.
func GetOneOrgHandler(c *gin.Context) {
	// Retrieve the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)

	// If there's an error retrieving the database connection, return an error.
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get a database connection",
		})
		return
	}

	// Get the organization name from the request parameter.
	orgNameParam := c.Param("orgName")

	// If the organization name is missing, return an error.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return
	}

	// Get the organization by name from the database.
	org, err := db.GetOrgByName(c, orgNameParam)

	// If there's an error retrieving the organization, return an error.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get the organization",
		})
		return
	}

	// Create a result struct with organization data to return.
	result := database.Org{
		Name:        org.Name,
		Description: org.Description,
		OwnerID:     org.OwnerID,
		CreatedAt:   org.CreatedAt,
		UpdatedAt:   org.UpdatedAt,
	}

	// Return the organization data.
	c.XML(http.StatusOK, result)
}

// UpdateOrgHandler updates an organization's information.
func UpdateOrgHandler(c *gin.Context) {
	// Extract token, database connection, and organization information.
	token, _, db := ExtractDBAndToken(c)

	// If the token is nil, return.
	if token == nil {
		return
	}

	// Verify ownership of the organization.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization ID is empty, return.
	if org.ID == "" {
		return
	}

	// Define a struct to hold XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML request to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Set the organization name and description based on the request or existing values.
	orgName := params.Name
	if orgName == "" {
		orgName = org.Name
	}

	orgDescription := params.Description
	if orgDescription == "" {
		orgDescription = org.Description
	}

	// Create arguments for updating the organization.
	args := database.UpdateOrgParams{
		ID:          org.ID,
		Name:        orgName,
		Description: orgDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the organization in the database.
	_, err := db.UpdateOrg(c, args)

	// If there's an error updating the organization, return an error.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to update the organization",
		})
		return
	}

	// Return the updated organization data.
	c.XML(http.StatusOK, args)
}

// DeleteOrgHandler deletes an organization.
func DeleteOrgHandler(c *gin.Context) {
	// Extract token, database connection, and organization information.
	token, _, db := ExtractDBAndToken(c)

	// If the token is nil, return.
	if token == nil {
		return
	}

	// Verify ownership of the organization.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization ID is empty, return.
	if org.ID == "" {
		return
	}

	// Delete the organization from the database.
	_, err := db.DeleteOrg(c, org.ID)

	// If there's an error deleting the organization, return an error.
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete the organization",
		})
		return
	}

	// Return a success message.
	c.XML(http.StatusOK, "Organization has been deleted successfully")
}

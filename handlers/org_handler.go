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
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

	orgID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate organization ID",
		})
		return
	}

	args := database.CreateOrgParams{
		ID:          orgID.String(),
		Name:        params.Name,
		Description: params.Description,
		CreatorID:   userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = db.CreateOrg(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new organization",
		})
		return
	}

	c.XML(http.StatusCreated, args)
}

func GetOrgsHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	orgs, err := db.GetAllOrgs(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organizations",
		})
		return
	}

	var orgMap []gin.H
	for _, org := range orgs {
		orgMap = append(orgMap, gin.H{
			"ID":          org.ID,
			"Name":        org.Name,
			"Description": org.Description,
			"CreatorID":   org.CreatorID,
			"CreatedAt":   org.CreatedAt,
			"UpdatedAt":   org.UpdatedAt,
		})
	}

	c.XML(http.StatusOK, gin.H{
		"Orgs": orgMap,
	})
}

func UpdateOrgHandler(c *gin.Context) {
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	org := VerifyOrgOwnership(c, token, db)
	if org.ID == "" {
		return
	}

	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	orgName := params.Name
	if orgName == "" {
		orgName = org.Name
	}

	orgDescription := params.Description
	if orgDescription == "" {
		orgDescription = org.Description
	}

	args := database.UpdateOrgParams{
		ID:          org.ID,
		Name:        orgName,
		Description: orgDescription,
		UpdatedAt:   time.Now(),
	}

	_, err := db.UpdateOrg(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to update organization",
		})
		return
	}

	c.XML(http.StatusOK, args)

}

func DeleteOrgHandler(c *gin.Context) {
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	org := VerifyOrgOwnership(c, token, db)
	if org.ID == "" {
		return
	}

	_, err := db.DeleteOrg(c, org.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete organization",
		})
		return
	}

	c.XML(http.StatusOK, "Organization has been deleted successfully")
}

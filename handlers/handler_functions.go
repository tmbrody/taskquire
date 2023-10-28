package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/tokenPackage"
)

func ExtractDBAndToken(c *gin.Context) (*jwt.Token, string, *database.Queries) {
	tokenString := tokenPackage.ExtractJWTTokenFromHeader(c.Request)
	if tokenString == "" {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "JWT token is missing or invalid",
		})
		return nil, "", nil
	}

	token, err := tokenPackage.ParseAndValidateJWTToken(tokenString)
	if err != nil {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Invalid JWT token",
		})
		return nil, "", nil
	}

	for _, tok := range tokenPackage.TokenMap {
		if tok.Claims.(jwt.MapClaims)["ID"] == token.Claims.(jwt.MapClaims)["ID"] {
			token = tok
			break
		}
	}

	if token.Claims.(jwt.MapClaims)["Revoked"] == true {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Trying to use revoked JWT token",
		})
		return nil, "", nil
	}

	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return nil, "", nil
	}

	return token, tokenString, db
}

func VerifyOrgOwnership(c *gin.Context, token *jwt.Token, db *database.Queries) database.Org {
	orgName := c.Query("orgName")
	if orgName == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return database.Org{}
	}

	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return database.Org{}
	}

	orgs, err := db.GetAllOrgs(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organizations",
		})
		return database.Org{}
	}

	var orgID string
	for _, org := range orgs {
		if org.Name == orgName {
			orgID = org.ID
			break
		}
	}
	if orgID == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization does not exist",
		})
		return database.Org{}
	}

	org, err := db.GetOrgByID(c, orgID)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return database.Org{}
	}

	if org.CreatorID != userID {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "You cannot modify this organization because you did not create it",
		})
		return database.Org{}
	}

	return org
}

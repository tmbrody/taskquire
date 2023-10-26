package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateOrgHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	var params struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	orgID, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to generate organization ID",
		})
		return
	}

	args := database.CreateOrgParams{
		ID:        orgID.String(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = db.CreateOrg(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new organization",
		})
		return
	}

	c.JSON(http.StatusCreated, args)
}

func GetOrgsHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	result, err := db.GetAllOrgs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organizations",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

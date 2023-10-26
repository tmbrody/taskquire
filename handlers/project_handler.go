package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateProjectHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	var params struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	projectID, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to generate project ID",
		})
		return
	}

	args := database.CreateProjectParams{
		ID:          projectID.String(),
		Name:        params.Name,
		Description: params.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = db.CreateProject(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new project",
		})
		return
	}

	c.JSON(http.StatusCreated, args)
}

func GetProjectsHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	result, err := db.GetAllProjects(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

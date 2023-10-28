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
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
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

	projectID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
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
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new project",
		})
		return
	}

	c.XML(http.StatusCreated, args)
}

func GetProjectsHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	projects, err := db.GetAllProjects(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return
	}

	var projectMap []gin.H
	for _, project := range projects {
		projectMap = append(projectMap, gin.H{
			"ID":          project.ID,
			"Name":        project.Name,
			"Description": project.Description,
			"CreatedAt":   project.CreatedAt,
			"UpdatedAt":   project.UpdatedAt,
		})
	}

	c.XML(http.StatusOK, gin.H{
		"Projects": projectMap,
	})
}

package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateTaskHandler(c *gin.Context) {
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
		ProjectID   string `json:"project_id"`
	}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	projects, err := db.GetAllProjects(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return
	}

	var projectID string
	for _, project := range projects {
		if params.ProjectID == project.ID {
			projectID = project.ID
			break
		}
	}

	if projectID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid project ID",
		})
		return
	}

	taskID, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to generate task ID",
		})
		return
	}

	var descriptionNullString sql.NullString

	if params.Description != "" {
		descriptionNullString.String = params.Description
		descriptionNullString.Valid = true
	} else {
		descriptionNullString.Valid = false
	}

	args := database.CreateTaskParams{
		ID:          taskID.String(),
		Name:        params.Name,
		Description: descriptionNullString,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ProjectID:   projectID,
	}

	_, err = db.CreateTask(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new task",
		})
		return
	}

	c.JSON(http.StatusCreated, args)
}

func GetTasksHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	result, err := db.GetAllTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get tasks",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

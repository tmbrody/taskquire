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
	// Get the project from the context.
	project, db := GetProjectByParamName(c)

	// If project has no name, return early without further processing.
	if project.Name == "" {
		return
	}

	// Define a struct to hold XML parameters and bind XML data to it.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML data from the request to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		// If there's an error in binding XML data, return a Bad Request error response.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Generate a new UUID for the task.
	taskID, err := uuid.NewUUID()
	if err != nil {
		// If there's an error generating the task ID, return an error response.
		c.XML(http.StatusBadRequest, gin.H{
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

	// Create the arguments for creating a new task.
	args := database.CreateTaskParams{
		ID:          taskID.String(),
		Name:        params.Name,
		Description: descriptionNullString,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ProjectID:   project.ID,
	}

	// Create a new task in the database.
	_, err = db.CreateTask(c, args)
	if err != nil {
		// If there's an error creating the task, return an error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new task",
		})
		return
	}

	// Return a successful response with the created task data.
	c.XML(http.StatusCreated, args)
}

func GetTasksHandler(c *gin.Context) {
	// Get the project from the context.
	project, db := GetProjectByParamName(c)

	// If project has no name, return early without further processing.
	if project.Name == "" {
		return
	}

	// Query the database to get all tasks.
	tasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		// If there was an error while querying tasks from the database, return an error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get tasks",
		})
		return
	}

	// Create a slice to hold the task data in a format suitable for XML serialization.
	var taskMap []gin.H
	for _, task := range tasks {
		// Append each task to the taskMap slice as a dictionary (map) with specific fields.
		taskMap = append(taskMap, gin.H{
			"ID":          task.ID,
			"Name":        task.Name,
			"Description": task.Description,
			"CreatedAt":   task.CreatedAt,
			"UpdatedAt":   task.UpdatedAt,
			"ProjectID":   task.ProjectID,
		})
	}

	// Return a successful response with the task data serialized as XML.
	c.XML(http.StatusOK, gin.H{
		"Tasks": taskMap,
	})
}

// GetOneTaskHandler is a handler function that retrieves a single task based on a name from the database and responds with XML data.
func GetOneTaskHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Get the task name from the parameter.
	taskNameParam := c.Param("taskName")

	// Check if the task name is missing. If so, return a Bad Request response.
	if taskNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Task name is missing",
		})
		return
	}

	// Retrieve a task with a certain name from the database
	task, err := db.GetTaskByName(c, taskNameParam)
	if err != nil {
		// If there is an error while getting the task, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get task",
		})
		return
	}

	// Prepare the result data to be sent back as an XML response
	result := database.Task{
		Name:        task.Name,
		Description: task.Description,
		ProjectID:   task.ProjectID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	// Return the public information of the task as an XML response
	c.XML(http.StatusOK, result)
}

func UpdateTaskHandler(c *gin.Context) {
	// Get the task from the context.
	task, db := GetTaskByParamName(c)

	// If task has no name, return early without further processing.
	if task.Name == "" {
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

	// Get the task name from the params or use the current task name if it's empty.
	taskName := params.Name
	if taskName == "" {
		taskName = task.Name
	}

	// Get the task description from the params or use the current task description if it's empty.
	taskDescriptionString := params.Description

	var taskDescription sql.NullString

	if taskDescriptionString != "" {
		taskDescription.String = taskDescriptionString
		taskDescription.Valid = true
	} else {
		taskDescription.String = params.Description
		taskDescription.Valid = false
	}

	// Prepare the arguments for updating the task in the database.
	args := database.UpdateTaskParams{
		ID:          task.ID,
		Name:        taskName,
		Description: taskDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the task in the database.
	_, err := db.UpdateTask(c, args)
	if err != nil {
		// If there is an error updating the task, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to update task",
		})
		return
	}

	// Respond with a success status and the updated task information.
	c.XML(http.StatusOK, args)
}

func DeleteTaskHandler(c *gin.Context) {
	// Get the project from the context.
	task, db := GetTaskByParamName(c)

	// If project has no name, return early without further processing.
	if task.Name == "" {
		return
	}

	// Attempt to delete the task from the database.
	_, err := db.DeleteTask(c, task.ID)
	if err != nil {
		// If there is an error deleting the task, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete task",
		})
		return
	}

	// Respond with a success status and a message indicating the task was deleted.
	c.XML(http.StatusOK, "Task has been deleted successfully")
}

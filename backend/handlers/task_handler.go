package handlers

import (
	"database/sql"
	"encoding/xml"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

// CreateTaskHandler is a HTTP request handler for creating a new task.
func CreateTaskHandler(c *gin.Context) {

	// Verify team membership and get project, team, database, and user ID.
	project, team, db, userID := VerifyTeamMembershipAndGetProject(c)

	// If team name is empty, return.
	if team.Name == "" {
		return
	}

	// Define a struct for holding XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Retrieve tasks for the project from the database.
	tasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get tasks",
		})
		return
	}

	// Check if a task with the same name already exists in the project.
	for _, task := range tasks {
		if task.Name == params.Name {
			c.XML(http.StatusBadRequest, config.ErrorResponse{
				Message: "Task name already exists",
			})
			return
		}
	}

	// Generate a new UUID for the task.
	taskID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Unable to generate task ID",
		})
		return
	}

	// Create a SQL null string for the task description.
	var descriptionNullString sql.NullString

	if params.Description != "" {
		descriptionNullString.String = params.Description
		descriptionNullString.Valid = true
	} else {
		descriptionNullString.Valid = false
	}

	// Initialize a parent task as an empty database.Task struct.
	parentTask := database.Task{}

	// Get the taskName parameter from the URL, if provided.
	taskNameParam := c.Param("taskName")

	// If taskNameParam is not empty, fetch the parent task by name from the database.
	if taskNameParam != "" {
		parentTask, err = db.GetTaskByName(c, taskNameParam)
		if err != nil {
			c.XML(http.StatusBadRequest, config.ErrorResponse{
				Message: "Parent task does not exist",
			})
			return
		}
	}

	// Define arguments for creating a new task.
	args := database.CreateTaskParams{
		ID:          taskID.String(),
		Name:        params.Name,
		Description: descriptionNullString,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ProjectID:   project.ID,
		TeamID:      team.ID,
		OwnerID:     userID,
		ParentID:    parentTask.ID,
	}

	// Create the new task in the database.
	_, err = db.CreateTask(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to create new task",
		})
		return
	}

	// Update timestamps for project, organization, and team.
	SetProjectUpdatedTime(c, db, project)
	SetOrgUpdatedTimeUsingProject(c, db, project)
	SetTeamUpdatedTime(c, db, team.ID)

	// Respond with the created task details.
	c.XML(http.StatusCreated, args)
}

// GetTasksHandler is a HTTP request handler for fetching tasks of a project.
func GetTasksHandler(c *gin.Context) {

	// Verify team membership and get project, team, and database.
	project, team, db, _ := VerifyTeamMembershipAndGetProject(c)

	// If team name is empty, return.
	if team.Name == "" {
		return
	}

	// Retrieve tasks for the project from the database.
	tasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get tasks",
		})
		return
	}

	// Create a slice to hold task information.
	var taskMap []gin.H
	for _, task := range tasks {
		taskMap = append(taskMap, gin.H{
			"ID":          task.ID,
			"Name":        task.Name,
			"Description": task.Description,
			"CreatedAt":   task.CreatedAt,
			"UpdatedAt":   task.UpdatedAt,
			"ProjectID":   task.ProjectID,
			"TeamID":      task.TeamID,
			"OwnerID":     task.OwnerID,
			"ParentID":    task.ParentID,
		})
	}

	xmlData, err := xml.Marshal(gin.H{"Tasks": taskMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "task")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetOneTaskHandler is a HTTP request handler for fetching a single task by name.
func GetOneTaskHandler(c *gin.Context) {

	// Get the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the taskName parameter from the URL.
	taskNameParam := c.Param("taskName")

	// If taskNameParam is empty, return an error.
	if taskNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Task name is missing",
		})
		return
	}

	// Fetch the task by name from the database.
	task, err := db.GetTaskByName(c, taskNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get task",
		})
		return
	}

	// Fetch subtasks for the task from the database.
	subtasks, err := db.GetSubtasksByParentID(c, task.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get subtasks",
		})
		return
	}

	// Create a slice to hold subtask names.
	var subtaskNames []string
	for _, subtask := range subtasks {
		subtaskNames = append(subtaskNames, subtask.Name)
	}

	// Create a TaskList structure for subtask names.
	subtaskList := TaskList{Tasks: subtaskNames}

	// Create a slice to hold task information.
	var taskMap []gin.H
	taskMap = append(taskMap, gin.H{
		"ID":          task.ID,
		"Name":        task.Name,
		"Description": task.Description,
		"ProjectID":   task.ProjectID,
		"TeamID":      task.TeamID,
		"CreatedAt":   task.CreatedAt,
		"UpdatedAt":   task.UpdatedAt,
		"ParentID":    task.ParentID,
		"Subtasks":    subtaskList,
	})

	// Respond with the task details.
	c.XML(http.StatusOK, taskMap)
}

// UpdateTaskHandler is a HTTP request handler for updating a task.
func UpdateTaskHandler(c *gin.Context) {

	// Get the task and database from the URL parameter.
	task, db := GetTaskByParamName(c)

	// If task name is empty, return.
	if task.Name == "" {
		return
	}

	// Define a struct for holding XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Determine the updated task name.
	taskName := params.Name
	if taskName == "" {
		taskName = task.Name
	}

	// Determine the updated task description.
	taskDescriptionString := params.Description
	var taskDescription sql.NullString

	if taskDescriptionString != "" {
		taskDescription.String = taskDescriptionString
		taskDescription.Valid = true
	} else {
		taskDescription.String = params.Description
		taskDescription.Valid = false
	}

	// Define arguments for updating the task.
	args := database.UpdateTaskParams{
		ID:          task.ID,
		Name:        taskName,
		Description: taskDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the task in the database.
	_, err := db.UpdateTask(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to update task",
		})
		return
	}

	// Update timestamps for the associated project.
	project := SetProjectUpdatedTimeUsingTask(c, db, task)

	// If project name is empty, return.
	if project.Name == "" {
		return
	}

	// Update timestamps for the organization using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Update timestamps for the team.
	SetTeamUpdatedTime(c, db, task.TeamID)

	// Respond with the updated task details.
	c.XML(http.StatusOK, args)
}

// DeleteTaskHandler is a HTTP request handler for deleting a task.
func DeleteTaskHandler(c *gin.Context) {

	// Get the task and database from the URL parameter.
	task, db := GetTaskByParamName(c)

	// If task name is empty, return.
	if task.Name == "" {
		return
	}

	// Fetch subtasks for the task from the database.
	subtasks, err := db.GetSubtasksByParentID(c, task.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get subtasks",
		})
		return
	}

	// Delete the parent task's subtasks from the database.
	if len(subtasks) > 0 {
		for _, subtask := range subtasks {
			_, err := db.DeleteTask(c, subtask.ID)
			if err != nil {
				c.XML(http.StatusInternalServerError, config.ErrorResponse{
					Message: "Unable to delete subtask",
				})
				return
			}
		}
	}

	// If the task has a parent task, update the parent task.
	if task.ParentID != "" {
		// Fetch the parent task from the database.
		parentTask, err := db.GetTaskByID(c, task.ParentID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get parent task",
			})
			return
		}

		// Define arguments for updating the task.
		args := database.UpdateTaskParams{
			ID:          parentTask.ID,
			Name:        parentTask.Name,
			Description: parentTask.Description,
			UpdatedAt:   time.Now(),
		}

		// Update the parent task in the database.
		_, err = db.UpdateTask(c, args)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to update parent task",
			})
			return
		}
	}

	// Delete the task from the database.
	_, err = db.DeleteTask(c, task.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to delete task",
		})
		return
	}

	// Update timestamps for the associated project.
	project := SetProjectUpdatedTimeUsingTask(c, db, task)

	// If project name is empty, return.
	if project.Name == "" {
		return
	}

	// Update timestamps for the organization using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Update timestamps for the team.
	SetTeamUpdatedTime(c, db, task.TeamID)

	// Respond with a success message.
	c.XML(http.StatusOK, "Task has been deleted successfully")
}

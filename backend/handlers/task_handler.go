package handlers

import (
	"bytes"
	"database/sql"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

// CreateTaskHandler handles the creation of a new task.
func CreateTaskHandler(c *gin.Context) {
	// Verify team membership and get project details
	project, team, db, userID := VerifyTeamMembershipAndGetProject(c)

	// If the team name is empty, return early
	if team.Name == "" {
		return
	}

	// Define the parameters for task creation from XML
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Parse the XML request body and bind it to the params struct
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Set task creation arguments
	args := SetTaskCreationArgs(c, db, project, team, userID, params.Name, params.Description)

	// Create a new task in the database
	_, err := db.CreateTask(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to create a new task",
		})
		return
	}

	// Update timestamps for project, organization, and team
	SetProjectUpdatedTime(c, db, project)
	SetOrgUpdatedTimeUsingProject(c, db, project)
	SetTeamUpdatedTime(c, db, team.ID)

	// Respond with the created task details
	c.XML(http.StatusCreated, args)
}

// GenerateTasksHandler generates tasks based on a project's existing tasks.
func GenerateTasksHandler(c *gin.Context) {
	// Verify team membership and get project details
	project, team, db, userID := VerifyTeamMembershipAndGetProject(c)

	// If the team name is empty, return early
	if team.Name == "" {
		return
	}

	// Retrieve existing project tasks from the database
	projectTasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get tasks",
		})
		return
	}

	// Filter existing tasks to get top-level tasks (tasks without a parent)
	existingTasks := []database.Task{}
	for _, projectTask := range projectTasks {
		if projectTask.ParentID == "" {
			existingTasks = append(existingTasks, projectTask)
		}
	}

	// Marshal existing tasks into XML format
	existingTasksXML, err := xml.Marshal(existingTasks)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	// Prepare the input data
	inputData := fmt.Sprintf("%s\n%s\n%s", project.Name, project.Description, string(existingTasksXML))

	// Execute a Python script to generate tasks based on existing ones
	cmd := exec.Command("python3", "chatgpt/generate_tasks.py")
	cmd.Stdin = strings.NewReader(inputData)

	// Capture the output of the Python script
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	// Run the Python script and check for errors
	err = cmd.Run()
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to generate tasks",
		})
		return
	}

	// Unmarshal the generated tasks from the Python script's output
	type Task struct {
		Name        string `xml:"name"`
		Description string `xml:"description"`
	}

	type Tasks struct {
		Tasks []Task `xml:"task"`
	}

	var tasks Tasks
	err = xml.Unmarshal(out.Bytes(), &tasks)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to unmarshal XML data",
		})
		return
	}

	// Create the generated tasks in the database and update timestamps
	args := database.CreateTaskParams{}
	for _, task := range tasks.Tasks {
		args = SetTaskCreationArgs(c, db, project, team, userID, task.Name, task.Description)

		if args == (database.CreateTaskParams{}) {
			return
		}

		_, err := db.CreateTask(c, args)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to create a new task",
			})
			return
		}

		// Update timestamps for project, organization, and team
		SetProjectUpdatedTime(c, db, project)
		SetOrgUpdatedTimeUsingProject(c, db, project)
		SetTeamUpdatedTime(c, db, team.ID)
	}

	// Respond with a success message
	c.XML(http.StatusOK, "Tasks have been generated successfully")
}

// GenerateSubtasksHandler generates subtasks for a given task.
func GenerateSubtasksHandler(c *gin.Context) {
	// Verify team membership and get project details
	project, team, db, userID := VerifyTeamMembershipAndGetProject(c)

	// If the team name is empty, return early
	if team.Name == "" {
		return
	}

	// Get the task details from the request
	task, _ := GetTaskByParamName(c)

	// If the task name is empty, return early
	if task.Name == "" {
		return
	}

	// Retrieve existing project tasks from the database
	projectTasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get tasks",
		})
		return
	}

	// Filter existing subtasks of the same parent task
	existingSubtasks := []database.Task{}
	for _, projectTask := range projectTasks {
		if projectTask.ParentID == task.ParentID && projectTask.ID != task.ID {
			existingSubtasks = append(existingSubtasks, projectTask)
		}
	}

	// Marshal existing subtasks into XML format
	existingSubtasksXML, err := xml.Marshal(existingSubtasks)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	// Prepare the input data
	inputData := fmt.Sprintf("%s\n%s\n%s", task.Name, task.Description.String, string(existingSubtasksXML))

	// Execute a Python script to generate subtasks based on existing ones
	cmd := exec.Command("python3", "chatgpt/generate_tasks.py")
	cmd.Stdin = strings.NewReader(inputData)

	// Capture the output of the Python script
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	// Run the Python script and check for errors
	err = cmd.Run()
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to generate subtasks",
		})
		return
	}

	// Unmarshal the generated subtasks from the Python script's output
	type Task struct {
		Name        string `xml:"name"`
		Description string `xml:"description"`
	}

	type Tasks struct {
		Tasks []Task `xml:"task"`
	}

	var tasks Tasks
	err = xml.Unmarshal(out.Bytes(), &tasks)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to unmarshal XML data",
		})
		return
	}

	// Create the generated subtasks in the database and update timestamps
	args := database.CreateTaskParams{}
	for _, task := range tasks.Tasks {
		args = SetTaskCreationArgs(c, db, project, team, userID, task.Name, task.Description)

		if args == (database.CreateTaskParams{}) {
			return
		}

		_, err := db.CreateTask(c, args)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to create a new subtask",
			})
			return
		}

		// Update timestamps for project, organization, and team
		SetProjectUpdatedTime(c, db, project)
		SetOrgUpdatedTimeUsingProject(c, db, project)
		SetTeamUpdatedTime(c, db, team.ID)
	}

	// Respond with a success message
	c.XML(http.StatusOK, "Subtasks have been generated successfully")
}

// GetTasksHandler retrieves and returns tasks for a project.
func GetTasksHandler(c *gin.Context) {
	// Verify team membership and get project details
	project, team, db, _ := VerifyTeamMembershipAndGetProject(c)

	// If the team name is empty, return early
	if team.Name == "" {
		return
	}

	// Retrieve all tasks for the project from the database
	allTasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get tasks",
		})
		return
	}

	// Filter top-level tasks (tasks without a parent)
	tasks := []database.Task{}
	for _, allTask := range allTasks {
		if allTask.ParentID == "" {
			tasks = append(tasks, allTask)
		}
	}

	// Prepare a map for task data to be converted to XML
	var taskMap []gin.H
	for _, task := range tasks {
		// Get subtasks for the current task
		subtasks, err := db.GetSubtasksByParentID(c, task.ID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get subtasks",
			})
			return
		}

		// Prepare subtask names
		var subtaskNames []string
		for _, subtask := range subtasks {
			subtaskNames = append(subtaskNames, subtask.Name)
		}

		// If no subtasks exist, add "None" to the list
		if subtaskNames == nil {
			subtaskNames = append(subtaskNames, "None")
		}

		// Create a task list for subtasks
		subtaskList := TaskList{Tasks: subtaskNames}

		// Create a map for the task data
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
			"Subtasks":    subtaskList,
		})
	}

	// Marshal the task data into XML
	xmlData, err := xml.Marshal(gin.H{"Tasks": taskMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	// Convert the XML data to a custom format
	xmlString, err := ConvertToCustomXML(xmlData, "task")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to a custom format",
		})
		return
	}

	// Respond with the custom XML data
	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetSubtasksHandler handles the HTTP request to retrieve subtasks of a parent task.
func GetSubtasksHandler(c *gin.Context) {
	// Verify team membership and get the project and database connection.
	project, team, db, _ := VerifyTeamMembershipAndGetProject(c)

	// Check if the team name is empty, and if so, return.
	if team.Name == "" {
		return
	}

	// Retrieve all tasks associated with the project.
	allTasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get tasks",
		})
		return
	}

	// Get the task name from the URL parameter.
	taskNameParam := c.Param("taskName")
	if taskNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Task name is missing",
		})
		return
	}

	// Retrieve the parent task by its name.
	parentTask, err := db.GetTaskByName(c, taskNameParam)
	if err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Parent task does not exist",
		})
		return
	}

	// Create a slice to store subtasks.
	tasks := []database.Task{}
	for _, allTasks := range allTasks {
		// Check if the task is a subtask of the parent task.
		if allTasks.ParentID == parentTask.ID {
			tasks = append(tasks, allTasks)
		}
	}

	// Create a list to store subtask names.
	var taskMap []gin.H
	for _, task := range tasks {
		// Retrieve subtasks for each task.
		subtasks, err := db.GetSubtasksByParentID(c, task.ID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get subtasks",
			})
			return
		}

		var subtaskNames []string
		for _, subtask := range subtasks {
			subtaskNames = append(subtaskNames, subtask.Name)
		}

		if subtaskNames == nil {
			subtaskNames = append(subtaskNames, "None")
		}

		subtaskList := TaskList{Tasks: subtaskNames}

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
			"Subtasks":    subtaskList,
		})
	}

	// Marshal the data to XML format.
	xmlData, err := xml.Marshal(gin.H{"Tasks": taskMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	// Convert the XML data to a custom format and send it as a response.
	xmlString, err := ConvertToCustomXML(xmlData, "task")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetOneTaskHandler handles the HTTP request to retrieve information about a single task.
func GetOneTaskHandler(c *gin.Context) {
	// Get the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the task name from the URL parameter.
	taskNameParam := c.Param("taskName")

	// Check if the task name is missing, and if so, return an error.
	if taskNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Task name is missing",
		})
		return
	}

	// Retrieve the task by its name.
	task, err := db.GetTaskByName(c, taskNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get task",
		})
		return
	}

	// Retrieve subtasks for the task.
	subtasks, err := db.GetSubtasksByParentID(c, task.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get subtasks",
		})
		return
	}

	// Create a list to store subtask names.
	var subtaskNames []string
	for _, subtask := range subtasks {
		subtaskNames = append(subtaskNames, subtask.Name)
	}

	// Create a subtask list.
	subtaskList := TaskList{Tasks: subtaskNames}

	// Create a map to store task information.
	var taskMap []gin.H
	taskMap = append(taskMap, gin.H{
		"ID":          task.ID,
		"Name":        task.Name,
		"Description": task.Description,
		"OwnerID":     task.OwnerID,
		"ProjectID":   task.ProjectID,
		"TeamID":      task.TeamID,
		"CreatedAt":   task.CreatedAt,
		"UpdatedAt":   task.UpdatedAt,
		"ParentID":    task.ParentID,
		"Subtasks":    subtaskList,
	})

	// Marshal the data to XML format.
	xmlData, err := xml.Marshal(gin.H{"Tasks": taskMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	// Convert the XML data to a custom format and send it as a response.
	xmlString, err := ConvertToCustomXML(xmlData, "task")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// UpdateTaskHandler handles the HTTP request to update a task's information.
func UpdateTaskHandler(c *gin.Context) {
	// Get the task and database connection.
	task, db := GetTaskByParamName(c)

	// Check if the task name is empty, and if so, return.
	if task.Name == "" {
		return
	}

	// Define a struct to hold the updated task parameters from XML.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Parse the XML request body and bind it to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Get the updated task name from the XML request or keep the existing name.
	taskName := params.Name
	if taskName == "" {
		taskName = task.Name
	}

	// Get the updated task description from the XML request.
	taskDescriptionString := params.Description
	var taskDescription sql.NullString

	if taskDescriptionString != "" {
		taskDescription.String = taskDescriptionString
		taskDescription.Valid = true
	} else {
		taskDescription.String = params.Description
		taskDescription.Valid = false
	}

	// Create arguments for updating the task.
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

	// Set the project's updated time using the task.
	project := SetProjectUpdatedTimeUsingTask(c, db, task)

	// Check if the project name is empty, and if so, return.
	if project.Name == "" {
		return
	}

	// Set the organization's updated time using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Set the team's updated time using the task's team ID.
	SetTeamUpdatedTime(c, db, task.TeamID)

	c.XML(http.StatusOK, args)
}

// DeleteTaskHandler handles the HTTP request to delete a task.
func DeleteTaskHandler(c *gin.Context) {
	// Get the task and database connection.
	task, db := GetTaskByParamName(c)

	// Check if the task name is empty, and if so, return.
	if task.Name == "" {
		return
	}

	// Retrieve subtasks associated with the task.
	subtasks, err := db.GetSubtasksByParentID(c, task.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get subtasks",
		})
		return
	}

	// Delete all subtasks if they exist.
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
		// Retrieve the parent task by its ID.
		parentTask, err := db.GetTaskByID(c, task.ParentID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get parent task",
			})
			return
		}

		// Create arguments for updating the parent task (no changes required).
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

	// Set the project's updated time using the task.
	project := SetProjectUpdatedTimeUsingTask(c, db, task)

	// Check if the project name is empty, and if so, return.
	if project.Name == "" {
		return
	}

	// Set the organization's updated time using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Set the team's updated time using the task's team ID.
	SetTeamUpdatedTime(c, db, task.TeamID)

	c.XML(http.StatusOK, "Task has been deleted successfully")
}

package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateProjectHandler(c *gin.Context) {
	// Get the database connection from the context.
	token, _, db := ExtractDBAndToken(c)

	// Verify that the user owns the organization based on the token.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization is not found or the user doesn't own it, return early.
	if org.ID == "" {
		return
	}

	// Define a structure to hold XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind XML request data to the params structure.
	if err := c.ShouldBindXML(&params); err != nil {
		// If the XML data is invalid, return a bad request error.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Generate a new project ID using the UUID library.
	projectID, err := uuid.NewUUID()
	if err != nil {
		// If unable to generate a project ID, return a bad request error.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate project ID",
		})
		return
	}

	// Prepare the arguments for creating a new project in the database.
	args := database.CreateProjectParams{
		ID:          projectID.String(),
		Name:        params.Name,
		Description: params.Description,
		OrgID:       org.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Create a new project in the database using the provided arguments.
	_, err = db.CreateProject(c, args)
	if err != nil {
		// If unable to create the project in the database, return an internal server error.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new project",
		})
		return
	}

	// Return a response indicating successful project creation.
	c.XML(http.StatusCreated, args)
}

func GetProjectsHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Get the organization name from the parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing. If so, return a Bad Request response.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return
	}

	org, err := db.GetOrgByName(c, orgNameParam)
	if err != nil {
		// If there was an error getting the organization from the database, return a 500 Internal Server Error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return
	}

	// Retrieve a list of projects from the database.
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		// If there was an error getting the projects from the database, return a 500 Internal Server Error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return
	}

	// Prepare a slice of project data in a suitable format for XML serialization.
	var projectMap []gin.H
	for _, project := range projects {
		projectMap = append(projectMap, gin.H{
			"ID":          project.ID,
			"Name":        project.Name,
			"Description": project.Description,
			"OrgID":       project.OrgID,
			"CreatedAt":   project.CreatedAt,
			"UpdatedAt":   project.UpdatedAt,
		})
	}

	// Return a 200 OK response with the list of projects in XML format.
	c.XML(http.StatusOK, gin.H{
		"Projects": projectMap,
	})
}

// GetOneProjectHandler is a handler function that retrieves a single project based on a name from the database and responds with XML data.
func GetOneProjectHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Get the organization name from the parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing. If so, return a Bad Request response.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Organization name is missing",
		})
		return
	}

	org, err := db.GetOrgByName(c, orgNameParam)
	if err != nil {
		// If there was an error getting the organization from the database, return a 500 Internal Server Error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organization",
		})
		return
	}

	// Retrieve a list of projects from the database.
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		// If there was an error getting the projects from the database, return a 500 Internal Server Error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get projects",
		})
		return
	}

	// Get the project name from the parameter.
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing. If so, return a Bad Request response.
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Project name is missing",
		})
		return
	}

	project := database.Project{}
	// Loop through the projects and check if the project name matches the name in the parameter.
	for _, p := range projects {
		if p.Name == projectNameParam {
			project = p
			break
		}
	}

	teams, err := db.GetAllTeamsFromProject(c, project.ID)
	if err != nil {
		// If there is an error while getting the team, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get project teams",
		})
		return
	}

	// get the names of all the users in the team
	var teamNames []string
	for _, team := range teams {
		t, err := db.GetTeamByID(c, team.TeamID)
		if err != nil {
			// If there is an error while getting the team, return an internal server error response
			c.XML(http.StatusInternalServerError, gin.H{
				"error": "Unable to get project's team name",
			})
			return
		}

		teamNames = append(teamNames, t.Name)
	}

	teamList := TeamList{Teams: teamNames}

	tasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		// If there is an error while getting the team, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get project teams",
		})
		return
	}

	// get the names of all the users in the team
	var taskNames []string
	for _, task := range tasks {
		taskNames = append(taskNames, task.Name)
	}

	taskList := TaskList{Tasks: taskNames}

	// Create a slice to hold the team data in a map format
	var teamMap []gin.H
	teamMap = append(teamMap, gin.H{
		"ID":          project.ID,
		"Name":        project.Name,
		"Description": project.Description,
		"OrgID":       project.OrgID,
		"CreatedAt":   project.CreatedAt,
		"UpdatedAt":   project.UpdatedAt,
		"Teams":       teamList,
		"Tasks":       taskList,
	})

	// Return the public information of the team as an XML response
	c.XML(http.StatusOK, teamMap)
}

func UpdateProjectHandler(c *gin.Context) {
	// Get the project from the context.
	project, db := VerifyOrgAndGetProjectByParamName(c)

	// If project has no name, return early without further processing.
	if project.Name == "" {
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

	// Get the project name from the params or use the current project name if it's empty.
	projectName := params.Name
	if projectName == "" {
		projectName = project.Name
	}

	// Get the project description from the params or use the current project description if it's empty.
	projectDescription := params.Description
	if projectDescription == "" {
		projectDescription = project.Description
	}

	// Prepare the arguments for updating the project in the database.
	args := database.UpdateProjectParams{
		ID:          project.ID,
		Name:        projectName,
		Description: projectDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the project in the database.
	_, err := db.UpdateProject(c, args)
	if err != nil {
		// If there is an error updating the project, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to update project",
		})
		return
	}

	// Respond with a success status and the updated project information.
	c.XML(http.StatusOK, args)
}

func AddTeamToProjectHandler(c *gin.Context) {
	// Extract the project, team and database connection from the context.
	project, team, db := GetTeamAndProjectByParam(c)

	// If the project is not found, return early.
	if project.ID == "" {
		return
	}

	// Prepare the arguments for updating the team in the database.
	args := database.AddTeamToProjectParams{
		ProjectID: project.ID,
		TeamID:    team.ID,
	}

	// Update the team in the database.
	_, err := db.AddTeamToProject(c, args)
	if err != nil {
		// If there is an error updating the team, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to add %s to %s", team.Name, project.Name),
		})
		return
	}

	// Respond with a success status and the updated team information.
	c.XML(http.StatusOK, fmt.Sprintf("%s has been added to %s", team.Name, project.Name))
}

func RemoveTeamFromProjectHandler(c *gin.Context) {
	// Extract the project, team and database connection from the context.
	project, team, db := GetTeamAndProjectByParam(c)

	// If the project is not found, return early.
	if project.ID == "" {
		return
	}

	args_get := database.GetOneProjectFromTeamParams{
		ProjectID: project.ID,
		TeamID:    team.ID,
	}

	_, err := db.GetOneProjectFromTeam(c, args_get)
	if err != nil {
		// If there is an error updating the team, respond with a Bad Request status and an error message.
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to find %s in %s", team.Name, project.Name),
		})
		return
	}

	// Prepare the arguments for updating the team in the database.
	args_remove := database.RemoveProjectFromTeamParams{
		ProjectID: project.ID,
		TeamID:    team.ID,
	}

	// Update the team in the database.
	_, err = db.RemoveProjectFromTeam(c, args_remove)
	if err != nil {
		// If there is an error updating the team, respond with an Internal Server Error status and an error message
		c.XML(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to remove %s from %s", team.Name, project.Name),
		})
		return
	}

	// Respond with a success status and the updated team information.
	c.XML(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s has been removed from %s", team.Name, project.Name),
	})
}

func DeleteProjectHandler(c *gin.Context) {
	// Get the project from the context.
	project, db := VerifyOrgAndGetProjectByParamName(c)

	// If project has no name, return early without further processing.
	if project.Name == "" {
		return
	}

	// Attempt to delete the project from the database.
	_, err := db.DeleteProject(c, project.ID)
	if err != nil {
		// If there is an error deleting the project, respond with an Internal Server Error status and an error message.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete project",
		})
		return
	}

	// Respond with a success status and a message indicating the project was deleted.
	c.XML(http.StatusOK, "Project has been deleted successfully")
}

package handlers

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

// CreateProjectHandler handles the creation of a new project.
func CreateProjectHandler(c *gin.Context) {
	// Extract database connection, token, and org information from the context.
	token, _, db := ExtractDBAndToken(c)

	// Verify ownership of the organization associated with the token.
	org := VerifyOrgOwnership(c, token, db)

	// If the organization ID is empty, return without creating a project.
	if org.ID == "" {
		return
	}

	// Define a struct to hold XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind the XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Generate a new UUID as the project ID.
	projectID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Unable to generate project ID",
		})
		return
	}

	// Create arguments for creating a new project in the database.
	args := database.CreateProjectParams{
		ID:          projectID.String(),
		Name:        params.Name,
		Description: params.Description,
		OrgID:       org.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Create the project in the database.
	_, err = db.CreateProject(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to create new project",
		})
		return
	}

	// Update the organization's updated time.
	SetOrgUpdatedTime(c, db, org)

	// Return the created project information.
	c.XML(http.StatusCreated, args)
}

// GetProjectsHandler retrieves a list of projects for a specific organization.
func GetProjectsHandler(c *gin.Context) {
	// Get the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the organization name from the URL parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Organization name is missing",
		})
		return
	}

	// Get the organization information from the database.
	org, err := db.GetOrgByName(c, orgNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get organization",
		})
		return
	}

	// Get the projects associated with the organization.
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return
	}

	// Create a list of project information to be returned.
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

	xmlData, err := xml.Marshal(gin.H{"Projects": projectMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "project")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// GetOneProjectHandler retrieves information for a specific project in an organization.
func GetOneProjectHandler(c *gin.Context) {
	// Get the database connection from the context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the organization name from the URL parameter.
	orgNameParam := c.Param("orgName")

	// Check if the organization name is missing.
	if orgNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Organization name is missing",
		})
		return
	}

	// Get the organization information from the database.
	org, err := db.GetOrgByName(c, orgNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get organization",
		})
		return
	}

	// Get the projects associated with the organization.
	projects, err := db.GetProjectsByOrgID(c, org.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get projects",
		})
		return
	}

	// Get the project name from the URL parameter.
	projectNameParam := c.Param("projectName")

	// Check if the project name is missing.
	if projectNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Project name is missing",
		})
		return
	}

	// Initialize an empty project structure.
	project := database.Project{}

	// Find the project with the specified name.
	for _, p := range projects {
		if p.Name == projectNameParam {
			project = p
			break
		}
	}

	// Get the teams associated with the project.
	teams, err := db.GetAllTeamsFromProject(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get project teams",
		})
		return
	}

	// Create a list of team names.
	var teamNames []string
	for _, team := range teams {
		t, err := db.GetTeamByID(c, team.TeamID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get project's team name",
			})
			return
		}
		teamNames = append(teamNames, t.Name)
	}

	// Create a structure to hold team names.
	teamList := TeamList{Teams: teamNames}

	// Get the tasks associated with the project.
	tasks, err := db.GetTasksByProjectID(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get project teams",
		})
		return
	}

	// Create a list of task names.
	var taskNames []string
	for _, task := range tasks {
		taskNames = append(taskNames, task.Name)
	}

	// Create a structure to hold task names.
	taskList := TaskList{Tasks: taskNames}

	// Create a list of project information to be returned.
	var projectMap []gin.H
	projectMap = append(projectMap, gin.H{
		"ID":          project.ID,
		"Name":        project.Name,
		"Description": project.Description,
		"OrgID":       project.OrgID,
		"CreatedAt":   project.CreatedAt,
		"UpdatedAt":   project.UpdatedAt,
		"Teams":       teamList,
		"Tasks":       taskList,
	})

	xmlData, err := xml.Marshal(gin.H{"Projects": projectMap})
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to marshal XML data",
		})
		return
	}

	xmlString, err := ConvertToCustomXML(xmlData, "project")
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to convert XML data to custom format",
		})
		return
	}

	c.Data(http.StatusOK, "application/xml", []byte(xmlString))
}

// UpdateProjectHandler updates the details of a project.
func UpdateProjectHandler(c *gin.Context) {
	// Verify the organization and get the project by parameter name.
	project, db := VerifyOrgAndGetProjectByParamName(c)

	// If the project name is empty, return without updating.
	if project.Name == "" {
		return
	}

	// Define a struct to hold XML request parameters.
	var params struct {
		Name        string `XML:"name"`
		Description string `XML:"description"`
	}

	// Bind the XML request body to the params struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Determine the updated project name.
	projectName := params.Name
	if projectName == "" {
		projectName = project.Name
	}

	// Determine the updated project description.
	projectDescription := params.Description
	if projectDescription == "" {
		projectDescription = project.Description
	}

	// Create arguments for updating the project in the database.
	args := database.UpdateProjectParams{
		ID:          project.ID,
		Name:        projectName,
		Description: projectDescription,
		UpdatedAt:   time.Now(),
	}

	// Update the project in the database.
	_, err := db.UpdateProject(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to update project",
		})
		return
	}

	// Update the organization's updated time.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Return the updated project information.
	c.XML(http.StatusOK, args)
}

// AddTeamToProjectHandler adds a team to a project.
func AddTeamToProjectHandler(c *gin.Context) {
	// Get the project, team, and database from the context.
	project, team, db := GetTeamAndProjectByParam(c)

	// If the project ID is empty, return without adding the team.
	if project.ID == "" {
		return
	}

	// Create arguments for adding a team to the project in the database.
	args := database.AddTeamToProjectParams{
		ProjectID: project.ID,
		TeamID:    team.ID,
	}

	// Add the team to the project in the database.
	_, err := db.AddTeamToProject(c, args)
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to add %s to %s", team.Name, project.Name),
		})
		return
	}

	// Update the organization's updated time using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Return a success message.
	c.XML(http.StatusOK, fmt.Sprintf("%s has been added to %s", team.Name, project.Name))
}

// RemoveTeamFromProjectHandler removes a team from a project.
func RemoveTeamFromProjectHandler(c *gin.Context) {
	// Get the project, team, and database from the context.
	project, team, db := GetTeamAndProjectByParam(c)

	// If the project ID is empty, return without removing the team.
	if project.ID == "" {
		return
	}

	// Create arguments for checking if the team is associated with the project.
	args_get := database.GetOneProjectFromTeamParams{
		ProjectID: project.ID,
		TeamID:    team.ID,
	}

	// Check if the team is associated with the project in the database.
	_, err := db.GetOneProjectFromTeam(c, args_get)
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Unable to find %s in %s", team.Name, project.Name),
		})
		return
	}

	// Create arguments for removing the team from the project in the database.
	args_remove := database.RemoveTeamFromProjectParams{
		ProjectID: project.ID,
		TeamID:    team.ID,
	}

	// Remove the team from the project in the database.
	_, err = db.RemoveTeamFromProject(c, args_remove)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to remove %s from %s", team.Name, project.Name),
		})
		return
	}

	// Update the organization's updated time using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Return a success message.
	c.XML(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s has been removed from %s", team.Name, project.Name),
	})
}

// DeleteProjectHandler deletes a project.
func DeleteProjectHandler(c *gin.Context) {
	// Verify the organization and get the project by parameter name.
	project, db := VerifyOrgAndGetProjectByParamName(c)

	// If the project name is empty, return without deleting.
	if project.Name == "" {
		return
	}

	// Delete the project from the database.
	_, err := db.DeleteProject(c, project.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to delete project",
		})
		return
	}

	// Update the organization's updated time using the project.
	SetOrgUpdatedTimeUsingProject(c, db, project)

	// Return a success message.
	c.XML(http.StatusOK, "Project has been deleted successfully")
}

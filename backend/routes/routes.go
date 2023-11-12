package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tmbrody/taskquire/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://taskquire.com", "https://www.taskquire.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.GET("/readiness", handlers.ReadinessHandler)
	r.GET("/error", handlers.ErrorHandler)

	r.GET("/", handlers.HomeHandler)
	r.GET("/about", handlers.AboutHandler)
}

func SetupApiRoutes(r *gin.RouterGroup) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://taskquire.com", "https://www.taskquire.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.POST("/orgs", handlers.CreateOrgHandler)
	r.GET("/orgs", handlers.GetOrgsHandler)
	r.GET("/orgs/:orgName", handlers.GetOneOrgHandler)
	r.PUT("/orgs/:orgName", handlers.UpdateOrgHandler)
	r.DELETE("/orgs/:orgName", handlers.DeleteOrgHandler)

	r.POST("/teams", handlers.CreateTeamHandler)
	r.GET("/teams", handlers.GetTeamsHandler)
	r.GET("/teams/:teamName", handlers.GetOneTeamHandler)
	r.PUT("/teams", handlers.AddUserToTeamHandler)
	r.PUT("/teams/:teamName", handlers.UpdateTeamHandler)
	r.DELETE("/teams", handlers.RemoveUserFromTeamHandler)
	r.DELETE("/teams/:teamName", handlers.DeleteTeamHandler)

	r.POST("/users", handlers.CreateUserHandler)
	r.GET("/users", handlers.GetUsersHandler)
	r.GET("/users/:userName", handlers.GetOneUserHandler)
	r.PUT("/users", handlers.UpdateUserHandler)
	r.DELETE("/users", handlers.DeleteUserHandler)

	r.POST("/login", handlers.LoginUserHandler)
	r.POST("/refresh", handlers.RefreshTokenHandler)
	r.POST("/revoke", handlers.RevokeTokenHandler)

	r.POST("/orgs/:orgName/projects", handlers.CreateProjectHandler)
	r.GET("/orgs/:orgName/projects", handlers.GetProjectsHandler)
	r.GET("/orgs/:orgName/projects/:projectName", handlers.GetOneProjectHandler)
	r.PUT("/orgs/:orgName/projects/:projectName", handlers.UpdateProjectHandler)
	r.DELETE("/orgs/:orgName/projects/:projectName", handlers.DeleteProjectHandler)

	r.GET("/orgs/:orgName/projects/:projectName/teams", handlers.GetTeamsFromProjectHandler)
	r.PUT("/orgs/:orgName/projects/:projectName/teams", handlers.AddTeamToProjectHandler)
	r.DELETE("/orgs/:orgName/projects/:projectName/teams", handlers.RemoveTeamFromProjectHandler)

	r.GET("/orgs/:orgName/projects/:projectName/teams/:teamName", handlers.GetOneTeamHandler)
	r.PUT("/orgs/:orgName/projects/:projectName/teams/:teamName", handlers.AddUserToTeamHandler)
	r.DELETE("/orgs/:orgName/projects/:projectName/teams/:teamName", handlers.RemoveUserFromTeamHandler)

	r.POST("/orgs/:orgName/projects/:projectName/teams/:teamName/tasks", handlers.CreateTaskHandler)
	r.GET("/orgs/:orgName/projects/:projectName/teams/:teamName/tasks", handlers.GetTasksHandler)

	r.POST("/orgs/:orgName/projects/:projectName/teams/:teamName/tasks/:taskName", handlers.CreateTaskHandler)
	r.GET("/orgs/:orgName/projects/:projectName/teams/:teamName/tasks/:taskName", handlers.GetOneTaskHandler)
	r.PUT("/orgs/:orgName/projects/:projectName/teams/:teamName/tasks/:taskName", handlers.UpdateTaskHandler)
	r.DELETE("/orgs/:orgName/projects/:projectName/teams/:teamName/tasks/:taskName", handlers.DeleteTaskHandler)
}

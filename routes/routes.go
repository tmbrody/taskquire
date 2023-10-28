package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tmbrody/taskquire/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*, http://*"},
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
	r.POST("/orgs", handlers.CreateOrgHandler)
	r.GET("/orgs", handlers.GetOrgsHandler)
	r.PUT("/orgs", handlers.UpdateOrgHandler)
	r.DELETE("/orgs", handlers.DeleteOrgHandler)

	r.POST("/teams", handlers.CreateTeamHandler)
	r.GET("/teams", handlers.GetTeamsHandler)

	r.POST("/users", handlers.CreateUserHandler)
	r.GET("/users", handlers.GetUsersHandler)
	r.PUT("/users", handlers.UpdateUserHandler)
	r.DELETE("/users", handlers.DeleteUserHandler)

	r.POST("/login", handlers.LoginUserHandler)
	r.POST("/refresh", handlers.RefreshTokenHandler)
	r.POST("/revoke", handlers.RevokeTokenHandler)

	r.POST("/projects", handlers.CreateProjectHandler)
	r.GET("/projects", handlers.GetProjectsHandler)

	r.POST("/tasks", handlers.CreateTaskHandler)
	r.GET("/tasks", handlers.GetTasksHandler)
}

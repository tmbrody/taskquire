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

	r.GET("/orgs", handlers.GetOrgsHandler)
	r.POST("/orgs", handlers.CreateOrgHandler)

	r.GET("/teams", handlers.GetTeamsHandler)
	r.POST("/teams", handlers.CreateTeamHandler)

	r.GET("/projects", handlers.GetProjectsHandler)
	r.POST("/projects", handlers.CreateProjectHandler)

	r.GET("/users", handlers.GetUsersHandler)
	r.POST("/users", handlers.CreateUserHandler)

	r.GET("/tasks", handlers.GetTasksHandler)
	r.POST("/tasks", handlers.CreateTaskHandler)
}

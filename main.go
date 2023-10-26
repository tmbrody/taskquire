package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/middleware"
	"github.com/tmbrody/taskquire/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	proxies := os.Getenv("TRUSTED_PROXIES")
	if proxies == "" {
		log.Fatal("TRUSTED_PROXIES environment variable is not set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	config.ApiCfg.DB = dbQueries

	r := gin.Default()

	trustedProxies := []string{proxies}
	err = r.SetTrustedProxies(trustedProxies)
	if err != nil {
		panic(err)
	}

	r.Use(middleware.CustomMiddleware())

	routes.SetupRoutes(r)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
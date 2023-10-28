package config

import "github.com/tmbrody/taskquire/internal/database"

type ApiConfig struct {
	DB        *database.Queries
	JwtSecret string
}

var ApiCfg ApiConfig

type ContextKey string

var DbContextKey ContextKey = "db"

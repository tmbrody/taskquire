package config

import "github.com/tmbrody/taskquire/internal/database"

type ApiConfig struct {
	DB  *database.Queries
	Org database.Org
}

var ApiCfg ApiConfig

type ContextKey string

var DbContextKey ContextKey = "db"

package config

import (
	"encoding/xml"

	"github.com/tmbrody/taskquire/internal/database"
)

type ApiConfig struct {
	DB        *database.Queries
	JwtSecret string
}

var ApiCfg ApiConfig

type ContextKey string

var DbContextKey ContextKey = "db"

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error"`
	Message string   `xml:"Message"`
}

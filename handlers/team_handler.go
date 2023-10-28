package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateTeamHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	var params struct {
		Name  string `XML:"name"`
		OrgID string `XML:"org_id"`
	}

	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	orgs, err := db.GetAllOrgs(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get organizations",
		})
		return
	}

	var orgID string
	for _, org := range orgs {
		if params.OrgID == org.ID {
			orgID = org.ID
			break
		}
	}

	if orgID == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid organization ID",
		})
		return
	}

	teamID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate team ID",
		})
		return
	}

	args := database.CreateTeamParams{
		ID:        teamID.String(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		OrgID:     orgID,
	}

	_, err = db.CreateTeam(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new team",
		})
		return
	}

	c.XML(http.StatusCreated, args)
}

func GetTeamsHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	teams, err := db.GetAllTeams(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get teams",
		})
		return
	}

	var teamMap []gin.H
	for _, team := range teams {
		teamMap = append(teamMap, gin.H{
			"ID":        team.ID,
			"Name":      team.Name,
			"CreatedAt": team.CreatedAt,
			"UpdatedAt": team.UpdatedAt,
			"OrgID":     team.OrgID,
		})
	}

	c.XML(http.StatusOK, gin.H{
		"Teams": teamMap,
	})
}

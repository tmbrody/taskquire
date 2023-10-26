package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
)

func CreateUserHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	var params struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to generate user ID",
		})
		return
	}

	args := database.CreateUserParams{
		ID:        userID.String(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  params.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = db.CreateUser(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new user",
		})
		return
	}

	c.JSON(http.StatusCreated, args)
}

func GetUsersHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	result, err := db.GetAllUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

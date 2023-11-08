package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/internal/database"
	"github.com/tmbrody/taskquire/tokenPackage"
	"golang.org/x/crypto/bcrypt"
)

// Constants for access and refresh token expiration durations
const AccessExpiration time.Duration = time.Hour
const RefreshExpiration time.Duration = 7 * (time.Hour * 24)

// CreateUserHandler handles the creation of a new user.
func CreateUserHandler(c *gin.Context) {
	// Retrieve the database connection from the Gin context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Parse XML request body into a struct
	var params struct {
		Name     string `XML:"name"`
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	// Check for XML parsing errors
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML.",
		})
		return
	}

	// Check if any of the fields are empty strings
	if CheckEmptyFields(params) {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid request payload: Missing or invalid fields.",
		})
		return
	}

	// Generate a unique user ID
	userID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Unable to generate user ID",
		})
		return
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to hash password",
		})
		return
	}

	// Create user data and insert it into the database
	args := database.CreateUserParams{
		ID:        userID.String(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.CreateUser(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to create new user",
		})
		return
	}

	// Prepare the response and send it
	result := database.User{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}
	c.XML(http.StatusCreated, result)
}

// GetUsersHandler retrieves a list of all users from the database.
func GetUsersHandler(c *gin.Context) {
	// Retrieve the database connection from the Gin context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Retrieve all users from the database
	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get users",
		})
		return
	}

	// Prepare the response with user data and send it
	var userMap []gin.H
	for _, user := range users {
		userMap = append(userMap, gin.H{
			"ID":        user.ID,
			"Name":      user.Name,
			"Email":     user.Email,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
		})
	}
	c.XML(http.StatusOK, gin.H{
		"Users": userMap,
	})
}

// GetOneUserHandler retrieves information about a specific user by username.
func GetOneUserHandler(c *gin.Context) {
	// Retrieve the database connection from the Gin context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Get the username from the request parameter
	userNameParam := c.Param("userName")

	// Check if the username is missing
	if userNameParam == "" {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "User name is missing",
		})
		return
	}

	// Retrieve user data by username from the database
	user, err := db.GetUserByName(c, userNameParam)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user",
		})
		return
	}

	// Retrieve teams associated with the user
	teams, err := db.GetAllTeamsByUser(c, user.ID)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get team users",
		})
		return
	}

	// Create a list of team names associated with the user
	var teamNames []string
	for _, team := range teams {
		t, err := db.GetTeamByID(c, team.TeamID)
		if err != nil {
			c.XML(http.StatusInternalServerError, config.ErrorResponse{
				Message: "Unable to get team user name",
			})
			return
		}
		teamNames = append(teamNames, t.Name)
	}

	// Prepare the response with user and team data and send it
	teamList := TeamList{Teams: teamNames}
	var teamMap []gin.H
	teamMap = append(teamMap, gin.H{
		"ID":        user.ID,
		"Name":      user.Name,
		"CreatedAt": user.CreatedAt,
		"UpdatedAt": user.UpdatedAt,
		"Teams":     teamList,
	})
	c.XML(http.StatusOK, teamMap)
}

// UpdateUserHandler updates user information.
func UpdateUserHandler(c *gin.Context) {
	// Extract the JWT token, database connection, and token package
	token, _, db := ExtractDBAndToken(c)

	// Check if there is no token
	if token == nil {
		return
	}

	// Retrieve the issuer from the token claims
	issuer := token.Claims.(jwt.MapClaims)["Issuer"].(string)

	// Check if the issuer is a refresh token
	if issuer == "taskquire-refresh" {
		c.XML(http.StatusUnauthorized, config.ErrorResponse{
			Message: "Using JWT refresh token when JWT access token is required",
		})
		return
	}

	// Parse XML request body into a struct
	var params struct {
		Name     string `XML:"name"`
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	// Check for XML parsing errors
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
		return
	}

	// Retrieve all users from the database
	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get users",
		})
		return
	}

	// Retrieve user ID from the token claims
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is empty
	if userID == "" {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user ID from JWT token",
		})
		return
	}

	// Initialize user information with default values
	userName := params.Name
	if userName == "" {
		for _, user := range users {
			if user.ID == userID {
				userName = user.Name
				break
			}
		}
	}

	userEmail := params.Email
	if userEmail == "" {
		for _, user := range users {
			if user.ID == userID {
				userEmail = user.Email
				break
			}
		}
	}

	userPassword := params.Password
	if userPassword == "" {
		for _, user := range users {
			if user.ID == userID {
				userPassword = user.Password
				break
			}
		}
	}

	// Hash the user's new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to hash password",
		})
		return
	}

	// Create user update parameters and update the user in the database
	args := database.UpdateUserParams{
		ID:        userID,
		Name:      userName,
		Email:     userEmail,
		Password:  string(hashedPassword),
		UpdatedAt: time.Now(),
	}
	_, err = db.UpdateUser(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Failed to update user data",
		})
		return
	}

	// Prepare the response with updated user data and send it
	result := database.User{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		UpdatedAt: args.UpdatedAt,
	}
	c.XML(http.StatusOK, result)
}

// DeleteUserHandler deletes a user account.
func DeleteUserHandler(c *gin.Context) {
	// Extract the JWT token, database connection, and token package
	token, _, db := ExtractDBAndToken(c)

	// Check if there is no token
	if token == nil {
		return
	}

	// Retrieve user ID from the token claims
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// Check if the user ID is empty
	if userID == "" {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get user ID from JWT token",
		})
		return
	}

	// Delete the user from the database
	_, err := db.DeleteUser(c, userID)

	// Check for deletion errors
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Failed to delete user",
		})
		return
	}

	// Revoke the user's tokens
	RevokeUserTokensHandler(c, userID)

	// Send a success message
	c.XML(http.StatusOK, "User account has been deleted successfully")
}

// LoginUserHandler handles user login and issues JWT tokens.
func LoginUserHandler(c *gin.Context) {
	// Retrieve the database connection from the Gin context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get database connection",
		})
		return
	}

	// Parse XML request body into a struct
	var params struct {
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	// Check for XML parsing errors
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, config.ErrorResponse{
			Message: "Invalid XML",
		})
	}

	// Retrieve all users from the database
	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, config.ErrorResponse{
			Message: "Unable to get users",
		})
		return
	}

	// Iterate through users to find a matching email
	for _, user := range users {
		if params.Email == user.Email {
			// Compare the hashed password with the provided password
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
				c.XML(http.StatusUnauthorized, config.ErrorResponse{
					Message: "Invalid password",
				})
				return
			}

			// Generate unique IDs for access and refresh tokens
			accessTokenID, err := uuid.NewUUID()
			if err != nil {
				c.XML(http.StatusInternalServerError, config.ErrorResponse{
					Message: "Unable to generate access token ID",
				})
				return
			}
			refreshTokenID, err := uuid.NewUUID()
			if err != nil {
				c.XML(http.StatusInternalServerError, config.ErrorResponse{
					Message: "Unable to generate refresh token ID",
				})
				return
			}

			// Create claims for access and refresh tokens
			accessClaims := jwt.MapClaims{
				"ID":        accessTokenID.String(),
				"Issuer":    "taskquire-access",
				"Subject":   user.ID,
				"IssuedAt":  jwt.NewNumericDate(time.Now()),
				"ExpiresAt": jwt.NewNumericDate(time.Now().Add(AccessExpiration)),
				"Revoked":   false,
			}
			refreshClaims := jwt.MapClaims{
				"ID":        refreshTokenID.String(),
				"Issuer":    "taskquire-refresh",
				"Subject":   user.ID,
				"IssuedAt":  jwt.NewNumericDate(time.Now()),
				"ExpiresAt": jwt.NewNumericDate(time.Now().Add(RefreshExpiration)),
				"Revoked":   false,
			}

			// Create access and refresh tokens
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

			// Add tokens to the token package
			tokenPackage.AddTokenToMap(token)
			tokenPackage.AddTokenToMap(refreshToken)

			// Sign access and refresh tokens
			signedToken, err := token.SignedString([]byte(config.ApiCfg.JwtSecret))
			if err != nil {
				c.XML(http.StatusInternalServerError, config.ErrorResponse{
					Message: "Unable to sign access token",
				})
				return
			}
			signedRefreshToken, err := refreshToken.SignedString([]byte(config.ApiCfg.JwtSecret))
			if err != nil {
				c.XML(http.StatusInternalServerError, config.ErrorResponse{
					Message: "Unable to sign refresh token",
				})
				return
			}

			// Prepare the response with user data and tokens
			c.XML(http.StatusOK, gin.H{
				"id":            accessClaims["Subject"],
				"name":          user.Name,
				"email":         user.Email,
				"access_token":  signedToken,
				"refresh_token": signedRefreshToken,
			})
			return
		}
	}

	// Send an error message if no user was found
	c.XML(http.StatusUnauthorized, config.ErrorResponse{
		Message: "User not found",
	})
}

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

const AccessExpiration time.Duration = time.Hour
const RefreshExpiration time.Duration = 7 * (time.Hour * 24)

// CreateUserHandler is an HTTP handler function that creates a new user in the database
func CreateUserHandler(c *gin.Context) {
	// Get a database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Define a struct to parse incoming XML data into
	var params struct {
		Name     string `XML:"name"`
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	// Parse the incoming XML request body into the 'params' struct
	if err := c.ShouldBindXML(&params); err != nil {
		// If there's an error parsing the XML, return a bad request error
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Generate a new UUID for the user ID
	userID, err := uuid.NewUUID()
	if err != nil {
		// If unable to generate a user ID, return a bad request error
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate user ID",
		})
		return
	}

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		// If unable to hash the password, return an internal server error
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to hash password",
		})
		return
	}

	// Prepare the user data to be inserted into the database
	args := database.CreateUserParams{
		ID:        userID.String(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert the user data into the database
	_, err = db.CreateUser(c, args)
	if err != nil {
		// If unable to create a new user in the database, return an internal server error
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new user",
		})
		return
	}

	// Prepare the result data to be sent back as XML response
	result := database.User{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}

	// Return the user data as a successful response with status 201 (Created)
	c.XML(http.StatusCreated, result)
}

// GetUsersHandler is an HTTP handler function that retrieves a list of users from the database
func GetUsersHandler(c *gin.Context) {
	// Attempt to retrieve the database connection from the Gin context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If there's an error retrieving the database connection, respond with an internal server error
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Retrieve the list of users from the database using the db connection
	users, err := db.GetAllUsers(c)
	if err != nil {
		// If there's an error retrieving users from the database, respond with an internal server error
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

	// Create a slice of gin.H maps to hold user data for XML serialization
	var userMap []gin.H
	for _, user := range users {
		// Populate the userMap with user data
		userMap = append(userMap, gin.H{
			"ID":        user.ID,
			"Name":      user.Name,
			"Email":     user.Email,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": user.UpdatedAt,
		})
	}

	// Respond with the list of users in XML format with a 200 OK status
	c.XML(http.StatusOK, gin.H{
		"Users": userMap,
	})
}

// GetOneUserHandler is a handler function that retrieves a single user based on a name from the database and responds with XML data.
func GetOneUserHandler(c *gin.Context) {
	// Get the database connection from the context
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Get the user name from the parameter.
	userNameParam := c.Param("userName")

	// Check if the user name is missing. If so, return a Bad Request response.
	if userNameParam == "" {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "User name is missing",
		})
		return
	}

	// Retrieve a user with a certain name from the database
	user, err := db.GetUserByName(c, userNameParam)
	if err != nil {
		// If there is an error while getting the user, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user",
		})
		return
	}

	teams, err := db.GetAllTeamsByUser(c, user.ID)
	if err != nil {
		// If there is an error while getting the team, return an internal server error response
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get team users",
		})
		return
	}

	// Get the names of all the teams for the user
	var teamNames []string
	for _, team := range teams {
		t, err := db.GetTeamByID(c, team.TeamID)
		if err != nil {
			// If there is an error while getting the team, return an internal server error response
			c.XML(http.StatusInternalServerError, gin.H{
				"error": "Unable to get team user name",
			})
			return
		}

		teamNames = append(teamNames, t.Name)
	}

	teamList := TeamList{Teams: teamNames}

	// Create a slice to hold the team data in a map format
	var teamMap []gin.H
	teamMap = append(teamMap, gin.H{
		"ID":        user.ID,
		"Name":      user.Name,
		"CreatedAt": user.CreatedAt,
		"UpdatedAt": user.UpdatedAt,
		"Teams":     teamList,
	})

	// Return the public information of the team as an XML response
	c.XML(http.StatusOK, teamMap)
}

// UpdateUserHandler is an HTTP handler function for updating user information.
func UpdateUserHandler(c *gin.Context) {
	// Extract the JWT token and database connection from the request context.
	token, _, db := ExtractDBAndToken(c)

	// If the token is nil, return early.
	if token == nil {
		return
	}

	// Get the issuer claim from the JWT token.
	issuer := token.Claims.(jwt.MapClaims)["Issuer"].(string)

	// Check if the issuer is a refresh token; if so, return an unauthorized error.
	if issuer == "taskquire-refresh" {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Using JWT refresh token when JWT access token is required",
		})
		return
	}

	// Define a struct to hold XML parameters from the request body.
	var params struct {
		Name     string `XML:"name"`
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	// Bind XML request body to the 'params' struct.
	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	// Retrieve a list of users from the database.
	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

	// Get the user ID from the JWT token's subject claim.
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// If the user ID is empty, return an error.
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

	// Initialize user information variables with default values from the database.
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

	// Hash the user's password using bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to hash password",
		})
		return
	}

	// Prepare arguments for updating user information in the database.
	args := database.UpdateUserParams{
		ID:        userID,
		Name:      userName,
		Email:     userEmail,
		Password:  string(hashedPassword),
		UpdatedAt: time.Now(),
	}

	// Update the user information in the database.
	_, err = db.UpdateUser(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user data",
		})
		return
	}

	// Create a user struct with updated information for response.
	result := database.User{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		UpdatedAt: args.UpdatedAt,
	}

	// Respond with the updated user information.
	c.XML(http.StatusOK, result)
}

func DeleteUserHandler(c *gin.Context) {
	// Extract the token, database connection, and user information from the context
	token, _, db := ExtractDBAndToken(c)

	// If the token is nil, there's no authenticated user, so return early
	if token == nil {
		return
	}

	// Extract the user ID from the JWT token's claims
	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)

	// If the extracted user ID is empty, return an error response
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

	// Attempt to delete the user from the database
	_, err := db.DeleteUser(c, userID)

	// If there was an error during the deletion process, return an error response
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	// Revoke any tokens associated with the deleted user
	RevokeUserTokensHandler(c, userID)

	// Return a success response indicating that the user account has been deleted
	c.XML(http.StatusOK, "User account has been deleted successfully")
}

// LoginUserHandler handles user login requests.
func LoginUserHandler(c *gin.Context) {
	// Attempt to retrieve the database connection from the Gin context.
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		// If unable to get the database connection, return an error response.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	// Define a struct to hold XML request parameters.
	var params struct {
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	// Attempt to bind XML request data into the 'params' struct.
	if err := c.ShouldBindXML(&params); err != nil {
		// If XML parsing fails, return a Bad Request error response.
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
	}

	// Retrieve a list of all users from the database.
	users, err := db.GetAllUsers(c)
	if err != nil {
		// If an error occurs while fetching users, return an internal server error.
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

	// Loop through the list of users to find a matching email.
	for _, user := range users {
		if params.Email == user.Email {
			// If the email matches, check the password.
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
				// If the password is invalid, return an Unauthorized error.
				c.XML(http.StatusUnauthorized, gin.H{
					"error": "Invalid password",
				})
				return
			}

			// Generate unique access and refresh token IDs.
			accessTokenID, err := uuid.NewUUID()
			if err != nil {
				c.XML(http.StatusInternalServerError, gin.H{
					"error": "Unable to generate access token ID",
				})
				return
			}

			refreshTokenID, err := uuid.NewUUID()
			if err != nil {
				c.XML(http.StatusInternalServerError, gin.H{
					"error": "Unable to generate refresh token ID",
				})
				return
			}

			// Define claims for access and refresh tokens.
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

			// Create JWT tokens with the specified claims.
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

			// Add the generated tokens to the token package.
			tokenPackage.AddTokenToMap(token)
			tokenPackage.AddTokenToMap(refreshToken)

			// Sign the access and refresh tokens with the JWT secret.
			signedToken, err := token.SignedString([]byte(config.ApiCfg.JwtSecret))
			if err != nil {
				// If token signing fails, return an internal server error.
				c.XML(http.StatusInternalServerError, gin.H{
					"error": "Unable to sign access token",
				})
				return
			}

			signedRefreshToken, err := refreshToken.SignedString([]byte(config.ApiCfg.JwtSecret))
			if err != nil {
				// If token signing fails, return an internal server error.
				c.XML(http.StatusInternalServerError, gin.H{
					"error": "Unable to sign refresh token",
				})
				return
			}

			// Return a successful response with user information and tokens.
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

	// If no matching user is found, return an Unauthorized error.
	c.XML(http.StatusUnauthorized, gin.H{
		"error": "User not found",
	})
}

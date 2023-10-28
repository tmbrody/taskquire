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

func CreateUserHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	var params struct {
		Name     string `XML:"name"`
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Unable to generate user ID",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to hash password",
		})
		return
	}

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
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to create new user",
		})
		return
	}

	result := database.User{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
	}

	c.XML(http.StatusCreated, result)
}

func GetUsersHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

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

func UpdateUserHandler(c *gin.Context) {
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	issuer := token.Claims.(jwt.MapClaims)["Issuer"].(string)
	if issuer == "taskquire-refresh" {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Using JWT refresh token when JWT access token is required",
		})
		return
	}

	var params struct {
		Name     string `XML:"name"`
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
		return
	}

	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to hash password",
		})
		return
	}

	args := database.UpdateUserParams{
		ID:        userID,
		Name:      userName,
		Email:     userEmail,
		Password:  string(hashedPassword),
		UpdatedAt: time.Now(),
	}

	_, err = db.UpdateUser(c, args)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user data",
		})
		return
	}

	result := database.User{
		ID:        args.ID,
		Name:      args.Name,
		Email:     args.Email,
		UpdatedAt: args.UpdatedAt,
	}

	c.XML(http.StatusOK, result)
}

func DeleteUserHandler(c *gin.Context) {
	token, _, db := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	userID := token.Claims.(jwt.MapClaims)["Subject"].(string)
	if userID == "" {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get user ID from JWT token",
		})
		return
	}

	_, err := db.DeleteUser(c, userID)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	RevokeUserTokensHandler(c, userID)

	c.XML(http.StatusOK, "User account has been deleted successfully")
}

func LoginUserHandler(c *gin.Context) {
	db, errBool := c.Value(string(config.DbContextKey)).(*database.Queries)
	if !errBool {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get database connection",
		})
		return
	}

	var params struct {
		Email    string `XML:"email"`
		Password string `XML:"password"`
	}

	if err := c.ShouldBindXML(&params); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error": "Invalid XML",
		})
	}

	users, err := db.GetAllUsers(c)
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to get users",
		})
		return
	}

	for _, user := range users {
		if params.Email == user.Email {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
				c.XML(http.StatusUnauthorized, gin.H{
					"error": "Invalid password",
				})
				return
			}

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

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

			tokenPackage.AddTokenToMap(token)
			tokenPackage.AddTokenToMap(refreshToken)

			signedToken, err := token.SignedString([]byte(config.ApiCfg.JwtSecret))
			if err != nil {
				c.XML(http.StatusInternalServerError, gin.H{
					"error": "Unable to sign access token",
				})
				return
			}

			signedRefreshToken, err := refreshToken.SignedString([]byte(config.ApiCfg.JwtSecret))
			if err != nil {
				c.XML(http.StatusInternalServerError, gin.H{
					"error": "Unable to sign refresh token",
				})
				return
			}

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

	c.XML(http.StatusUnauthorized, gin.H{
		"error": "User not found",
	})
}

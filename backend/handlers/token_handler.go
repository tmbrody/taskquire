package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tmbrody/taskquire/config"
	"github.com/tmbrody/taskquire/tokenPackage"
)

func RefreshTokenHandler(c *gin.Context) {
	// Extract the database connection and refreshToken from the context
	refreshToken, _, _ := ExtractDBAndToken(c)

	// If no token is found, return early
	if refreshToken == nil {
		return
	}

	// Get the issuer from the token's claims
	issuer := refreshToken.Claims.(jwt.MapClaims)["Issuer"].(string)

	// Check if the issuer is "taskquire-access"
	if issuer == "taskquire-access" {
		// If the issuer is incorrect, return an unauthorized response with an error message
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Using JWT access token when JWT refresh token is required",
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

	subject := refreshToken.Claims.(jwt.MapClaims)["Subject"].(string)

	// Create a new access token with the same subject as the refresh token.
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       accessTokenID.String(),
		"Subject":  subject,
		"Issuer":   "taskquire-access",
		"IssuedAt": jwt.NewNumericDate(time.Now()),
		// "ExpiresAt": jwt.NewNumericDate(time.Now().Add(AccessExpiration)),
		"Revoked": false,
	})

	// Sign the access token with the JWT secret.
	signedToken, err := accessToken.SignedString([]byte(config.ApiCfg.JwtSecret))
	if err != nil {
		c.XML(http.StatusInternalServerError, gin.H{
			"error": "Unable to sign JWT access token",
		})
		return
	}

	// Display the access token in the response.
	c.XML(http.StatusOK, gin.H{
		"access_token": signedToken,
	})
}

func RevokeTokenHandler(c *gin.Context) {
	token, _, _ := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	token.Claims.(jwt.MapClaims)["Revoked"] = true

	c.XML(http.StatusOK, "JWT token successfully revoked")
}

func RevokeUserTokensHandler(c *gin.Context, userId string) {
	for _, token := range tokenPackage.TokenMap {
		if token.Claims.(jwt.MapClaims)["Subject"].(string) == userId {
			token.Claims.(jwt.MapClaims)["Revoked"] = true
		}
	}
}

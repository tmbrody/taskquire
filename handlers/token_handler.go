package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tmbrody/taskquire/tokenPackage"
)

func RefreshTokenHandler(c *gin.Context) {
	token, _, _ := ExtractDBAndToken(c)

	if token == nil {
		return
	}

	issuer := token.Claims.(jwt.MapClaims)["Issuer"].(string)
	if issuer == "taskquire-access" {
		c.XML(http.StatusUnauthorized, gin.H{
			"error": "Using JWT access token when JWT refresh token is required",
		})
		return
	}

	for _, tok := range tokenPackage.TokenMap {
		if tok.Claims.(jwt.MapClaims)["Subject"] == token.Claims.(jwt.MapClaims)["Subject"] {
			if tok.Claims.(jwt.MapClaims)["Revoked"] == true {
				c.XML(http.StatusUnauthorized, gin.H{
					"error": "JWT access token is already revoked",
				})
				return
			}

			tok.Claims.(jwt.MapClaims)["ExpiresAt"] = jwt.NewNumericDate(time.Now().Add(AccessExpiration))
			c.XML(http.StatusOK, "JWT access token has been successfully refreshed")
			return
		}
	}

	c.XML(http.StatusBadRequest, "No JWT access tokens found for this refeesh token")
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

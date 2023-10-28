package tokenPackage

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tmbrody/taskquire/config"
)

var TokenMap = make(map[string]*jwt.Token)

func AddTokenToMap(token *jwt.Token) {
	id := token.Claims.(jwt.MapClaims)["ID"].(string)
	TokenMap[id] = token
}

func ExtractJWTTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	if len(authHeader) > 7 && strings.Contains(authHeader, " ") {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 {
			return parts[1]
		}
	}
	return ""
}

func ParseAndValidateJWTToken(tokenString string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, errors.New("no token provided")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.ApiCfg.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

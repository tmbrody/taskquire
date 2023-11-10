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

// ExtractJWTTokenFromHeader extracts a JWT token from an HTTP request's Authorization header.
// If the header is not present or doesn't contain a valid token, it returns an empty string.
func ExtractJWTTokenFromHeader(r *http.Request) string {
	// Get the value of the "Authorization" header from the request.
	authHeader := r.Header.Get("Authorization")

	// If the Authorization header is empty, return an empty string.
	if authHeader == "" {
		return ""
	}

	// Check if the Authorization header has at least 8 characters and contains a space.
	if len(authHeader) > 7 && strings.Contains(authHeader, " ") {
		// Split the Authorization header into two parts based on the space character.
		parts := strings.SplitN(authHeader, " ", 2)

		// If there are exactly two parts after splitting, return the second part,
		// which should be the JWT token.
		if len(parts) == 2 {
			return parts[1]
		}
	}

	// If the Authorization header does not meet the criteria or splitting fails,
	// return an empty string.
	return ""
}

// ParseAndValidateJWTToken takes a JWT token string as input and parses it,
// then validates its signature using a predefined HMAC secret key.
// It returns the parsed token if successful, or an error if any validation fails.
func ParseAndValidateJWTToken(tokenString string) (*jwt.Token, error) {
	// Check if the tokenString is empty.
	if tokenString == "" {
		// If no token is provided, return an error.
		return nil, errors.New("no token provided")
	}

	// Parse the JWT token using the provided tokenString and a custom validation function.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the token uses an HMAC signing method.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// If the signing method is not HMAC, return an error indicating an invalid signing method.
			return nil, errors.New("invalid signing method")
		}

		// Return the HMAC secret key for validation from a configuration file.
		return []byte(config.ApiCfg.JwtSecret), nil
	})

	// Check for any errors during token parsing or validation.
	if err != nil {
		// If there was an error, return it.
		return nil, err
	}

	// Return the parsed and validated token.
	return token, nil
}

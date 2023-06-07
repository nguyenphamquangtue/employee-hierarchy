package middleware

import (
	"employee-hierarchy-api/external/dto"
	"employee-hierarchy-api/internal/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"os"
	"strings"
	"time"
)

// GenerateJWT This method generates a JWT
func GenerateJWT(user *dto.User) (string, error) {
	// Create a claims object with user information
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // JWT expiration time
	}

	// Create a JWT with claims and a secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("SECRET_KEY")) // Secret key, replace with an actual and secure key

	// Sign and convert the JWT to a string
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Lấy AccessToken từ header hoặc query parameter
		accessToken := c.Request().Header.Get("Authorization")
		if accessToken == "" {
			accessToken = c.QueryParam("access_token")
		}
		token := strings.Split(accessToken, " ")
		if len(token) > 2 || len(token) < 2 {
			return response.R401(c, echo.Map{}, "Invalid Authorization header")
		}
		if !IsValidAccessToken(token[1]) {
			return response.R401(c, echo.Map{}, "Invalid access token")
		}

		return next(c)
	}
}

func IsValidAccessToken(accessToken string) bool {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return false
	}

	return true
}

package middleware

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var tokenBlacklist = make(map[string]time.Time)

func Revoke(tokenString string) error {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the key to validate the token's signature
		// You can use a shared secret key or a public key depending on your token signing mechanism
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to parse token: %v", err))
	}

	// Check if the token is expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(expirationTime) {
			return errors.New("token has expired")
		}
	} else {
		return fmt.Errorf("Invalid token")
	}

	// Add the token to the blacklist
	tokenBlacklist[tokenString] = time.Now()
	fmt.Println("Token revoked successfully")
	return nil
}

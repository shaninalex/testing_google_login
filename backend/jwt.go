package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(id, email string) (string, error) {
	// tokenValidityDuration := 24 * time.Hour
	// now := time.Now()
	// exp := now.Add(tokenValidityDuration)
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour)
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

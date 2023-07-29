package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(id, email string) (string, error) {
	// https://www.golinuxcloud.com/golang-jwt/
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func validateToken(t string) (string, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		log.Println(err)
		log.Println("error on parsing")
		return "", err
	}

	if token == nil {
		log.Println(err)
		log.Println("invalid token")
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("couldn't parse claims")
		return "", err
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		log.Println("token expired")
		return "", err
	}

	return claims["sub"].(string), err
}

// if token == nil {
// 	fmt.Fprintf(w, "invalid token")
// 	return errors.New("Token error")
// }

// claims, ok := token.Claims.(jwt.MapClaims)
// if !ok {
// 	fmt.Fprintf(w, "couldn't parse claims")
// 	return errors.New("Token error")
// }

// exp := claims["exp"].(float64)
// if int64(exp) < time.Now().Local().Unix() {
// 	fmt.Fprintf(w, "token expired")
// 	return errors.New("Token error")
// }

// return nil

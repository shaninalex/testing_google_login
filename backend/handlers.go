package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) handleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.JSON(http.StatusOK, gin.H{"link": url})
}

func (app *App) handleGoogleCallback(c *gin.Context) {
	state := c.DefaultQuery("state", "")
	if state != oauthStateString {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.DefaultQuery("code", "")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("OAuth2 exchange error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Fetch the user's profile information
	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Error fetching user info: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Printf("Error decoding user info: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	name := userInfo["name"].(string)
	email := userInfo["email"].(string)
	avatar := userInfo["picture"].(string)
	id, err := app.database.CreateSocialUser(name, email, avatar) // , "google"
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant create user or user already exists"})
		return
	}

	jwt, err := CreateJWT(id, email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant create token"})
		return
	}
	// TODO: Generate auth cookie instead of using Google OAuth access token
	// 		 because we mait have regular email/password login flow
	c.SetCookie("token", jwt, 3600, "/", "localhost", true, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200/auth/profile")
}

func (app *App) handleUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User profile endpoint",
	})
}

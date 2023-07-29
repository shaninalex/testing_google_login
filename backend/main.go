package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random" // Replace this with a random string of your choice
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		// RedirectURL:  "http://localhost:4200/auth/google/callback",
		RedirectURL: "http://localhost:8080/api/v1/auth/google/callback",
		Scopes:      []string{"openid", "email", "profile"},
		Endpoint:    google.Endpoint,
	}
}

func main() {
	r := gin.Default()

	r.GET("/api/v1/auth/google/login", handleLogin)
	r.GET("/api/v1/auth/google/callback", handleCallback)
	r.GET("/api/v1/user/profile", handleUserProfile)

	r.Run(":8080")
}

func handleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	// c.Redirect(http.StatusTemporaryRedirect, url)
	c.JSON(http.StatusOK, gin.H{"link": url})
}

func handleCallback(c *gin.Context) {
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

	fmt.Println("Access Token:", token.AccessToken)
	fmt.Println("Token Type:", token.TokenType)
	fmt.Println("Expiry:", token.Expiry.Format("2006-01-02 15:04:05"))

	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Avatar:", avatar)

	// TODO: Generate auth cookie instead of using Google OAuth access token
	// 		 because we mait have regular email/password login flow
	c.SetCookie("token", token.AccessToken, 3600, "/", "localhost", true, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:4200/auth/profile")
}

func handleUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User profile endpoint",
	})
}

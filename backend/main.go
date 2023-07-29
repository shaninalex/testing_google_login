package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	CLIENT_ID         = os.Getenv("CLIENT_ID")
	CLIENT_SECRET     = os.Getenv("CLIENT_SECRET")
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random" // Replace this with a random string of your choice
)

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     CLIENT_ID,
		ClientSecret: CLIENT_SECRET,
		RedirectURL:  "http://localhost:4200/auth/callback", // add this in google auth client id (https://console.cloud.google.com/apis/credentials)
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
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
	c.Redirect(http.StatusTemporaryRedirect, url)
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

	// Normally, you'd save the token for later use (e.g., to make authenticated requests to Google APIs)
	// For this example, we'll just print the token.
	fmt.Println("Access Token:", token.AccessToken)
	fmt.Println("Refresh Token:", token.RefreshToken)
	fmt.Println("Token Type:", token.TokenType)
	fmt.Println("Expiry:", token.Expiry.Format("2006-01-02 15:04:05"))

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully authenticated with Google!",
	})
}

func handleUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User profile endpoint",
	})
}

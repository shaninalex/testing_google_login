package main

import (
	"backend/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "random" // Replace this with a random string of your choice
)

func init() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		// RedirectURL:  "http://localhost:4200/auth/google/callback",
		RedirectURL: "http://localhost:8080/api/v1/auth/google/callback",
		Scopes:      []string{"openid", "email", "profile"},
		Endpoint:    google.Endpoint,
	}
}

type App struct {
	router   *gin.Engine
	database *database.Database
}

func main() {
	db, err := database.CreateConnection(os.Getenv("DATABASE_URI"))
	if err != nil {
		log.Fatal("Can't connect to database")
	}

	app := &App{
		router:   gin.Default(),
		database: db,
	}

	app.router.GET("/api/v1/auth/google/login", app.handleGoogleLogin)
	app.router.GET("/api/v1/auth/google/callback", app.handleGoogleCallback)
	app.router.POST("/api/v1/auth/login", app.handleRegularLogin)
	app.router.POST("/api/v1/auth/register", app.handleRegularRegister)
	app.router.GET("/api/v1/user/profile", app.handleUserProfile)
	app.router.Run(":8080")
}

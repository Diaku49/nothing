package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/Diaku49/nothing.git/db"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type AppHandler struct {
	DB                *gorm.DB
	CSRF              string
	GoogleOauthConfig *oauth2.Config
}

func NewAppHandler() (*AppHandler, error) {
	app := &AppHandler{}

	// Setup Database
	db, err := db.InitDb()
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	app.DB = db

	// Setup CSRF
	csrf := os.Getenv("CSRF_PROTECTION")
	app.CSRF = csrf

	// Google OauthConfig
	app.GoogleOauthConfig = GoogleOauthConfig()

	return app, nil
}

func GoogleOauthConfig() *oauth2.Config {
	googleAppClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URL")

	if googleAppClientID == "" || googleClientSecret == "" || redirectURL == "" {
		log.Fatal("FATAL: Ensure GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_REDIRECT_URL, and APP_JWT_SECRET env vars are set.")
	}

	googleOauthConfig := &oauth2.Config{
		RedirectURL:  redirectURL,
		ClientID:     googleAppClientID,
		ClientSecret: googleClientSecret,
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return googleOauthConfig
}

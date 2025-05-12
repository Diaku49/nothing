package handlers

import (
	"fmt"
	"net/http"

	"github.com/Diaku49/nothing.git/internals/helpers/jwt"
	hOauth "github.com/Diaku49/nothing.git/internals/helpers/oauth"
	hUser "github.com/Diaku49/nothing.git/internals/helpers/user"
	m "github.com/Diaku49/nothing.git/models"
	"github.com/gofiber/fiber/v2"
)

func (ah *AppHandler) GoogleLogin(c *fiber.Ctx) error {
	state := hOauth.GenerateRandomStr()

	// Set cookie for CSRF attacks
	c.Cookie(&fiber.Cookie{
		Name:     ah.CSRF,
		Value:    state,
		Path:     "/",
		MaxAge:   600,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
	})

	url := ah.GoogleOauthConfig.AuthCodeURL(state)
	return c.Redirect(url, http.StatusTemporaryRedirect)
}

func (ah *AppHandler) GoogleCallBack(c *fiber.Ctx) error {
	userCtx := c.UserContext()
	// Check our state
	stateFromCookie := c.Cookies(ah.CSRF)
	if stateFromCookie == "" {
		fmt.Println("State cookie is missing.")
		return c.Status(http.StatusBadRequest).JSON(
			m.ResponseHttp{
				Data:    nil,
				Success: false,
				Message: "Missing state cookie",
			})
	}
	c.ClearCookie(ah.CSRF)

	// Checking oauth code
	googleAuthCode := c.Query("code")
	if googleAuthCode == "" {
		fmt.Println("Google auth code missing from callback")
		errorReason := c.Query("error", "unknown_error")
		return c.Status(fiber.StatusUnauthorized).JSON(m.ResponseHttp{
			Success: false,
			Message: "Google authorization failed",
			Data:    errorReason,
		})
	}

	// Excahnge Google code for Token
	googleToken, err := ah.GoogleOauthConfig.Exchange(userCtx, googleAuthCode)
	if err != nil {
		fmt.Printf("Failed to exchange Google code: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "Failed to excahnge Google code",
				Data:    err,
			})
	}

	// Validate Token
	validatedGoogleToken, err := hOauth.ValidateGoogleToken(googleToken, &userCtx, ah.GoogleOauthConfig.ClientID)
	if err != nil {
		fmt.Printf("Invalid google Token: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "Invalid google Token",
				Data:    err,
			})
	}

	// Extract user info
	googleUserID := validatedGoogleToken.Subject
	email := validatedGoogleToken.Claims["email"].(string)
	name := validatedGoogleToken.Claims["name"].(string)

	// Check user existense
	UserExist, err := hUser.UserExistByEmail(email, ah.DB)
	if err != nil {
		fmt.Printf("DB error: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "db error",
				Data:    err,
			})
	}
	if UserExist {
		fmt.Println("This user already exist")
		return c.Status(http.StatusBadRequest).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "user with this email already exist",
			})
	}

	// Generate RefreshToken
	refreshTokenClaims := jwt.RefreshTokenClaims{
		GoogleID: googleUserID,
	}
	refreshToken, err := jwt.CreateRefreshToken(refreshTokenClaims)
	if err != nil {
		fmt.Println("This user already exist")
		return c.Status(http.StatusInternalServerError).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "failed creating token",
				Data:    err,
			})
	}

	// Create new user and store refresh token
	user := m.User{FullName: name, Email: email, RefreshTokenoh: &refreshToken}
	result := ah.DB.Create(&user)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "failed creating user",
				Data:    result.Error,
			})
	}

	accessTokenClaims := jwt.AccessTokenClaims{
		ID:       user.ID,
		Email:    email,
		GoogleID: googleUserID,
	}
	accessToken, err := jwt.CreateAccessToken(accessTokenClaims)
	if err != nil {
		fmt.Println("This user already exist")
		return c.Status(http.StatusInternalServerError).JSON(
			m.ResponseHttp{
				Success: false,
				Message: "failed creating token",
				Data:    err,
			})
	}

	// Send AccessToken to client
	return c.Status(http.StatusCreated).JSON(
		m.ResponseHttp{
			Success: true,
			Message: "Succesfully signed up",
			Data:    accessToken,
		})
}

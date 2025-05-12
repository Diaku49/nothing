package routes

import (
	ah "github.com/Diaku49/nothing.git/internals/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Routes(ah *ah.AppHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	}))

	v1 := app.Group("/v1")

	// AuthRoutes
	authRoute := v1.Group("/auth")
	authRoute.Post("/signup", ah.Signup)
	authRoute.Post("/login", ah.Login)
	authRoute.Get("/google/login", ah.GoogleLogin)
	authRoute.Get("/google/callback", ah.GoogleCallBack)

	// OtherRoutes

	return app
}

package cmd

import (
	"os"

	ah "github.com/Diaku49/nothing.git/internals/handlers"
	"github.com/Diaku49/nothing.git/internals/routes"
	"github.com/joho/godotenv"
)

func StartServer() {
	// Env Setup
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Creating App
	a, err := ah.NewAppHandler()
	if err != nil {
		panic(err)
	}

	app := routes.Routes(a)

	app.Listen(":" + port)
}

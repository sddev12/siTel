package main

import (
	"log"
	"todo-iam/app"
	"todo-iam/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Set up server
	app, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	// Routes
	app.Echo.POST("/register", func(c echo.Context) error {
		return handlers.RegisterHandler(app, c)
	})

	app.Echo.POST("/login", func(c echo.Context) error {
		return handlers.LoginHandler(app, c)
	})

	app.Echo.Logger.Fatal((app.Echo.Start(":3002")))
}

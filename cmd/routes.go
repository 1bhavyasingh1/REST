package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/1bhavyasingh1/REST/handlers"
)


func setupRoutes(app *fiber.App){
	app.Get("/", handlers.Facts)

    app.Post("/fact", handlers.CreateFact)
}
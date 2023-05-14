package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/1bhavyasingh1/REST/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    

    app.Listen(":3000")
}
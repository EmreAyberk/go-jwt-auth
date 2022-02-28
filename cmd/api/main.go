package main

import (
	"github.com/EmreAyberk/go-jwt-auth/pkg/database"
	"github.com/EmreAyberk/go-jwt-auth/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {

	_, err := database.Connect()
	if err != nil {
		panic("could not connect database")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	log.Fatalln(app.Listen(":3000"))
}

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/api"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	fmt.Println("Server Started")

	// Health check route
	app.Get("/ping", func(c *fiber.Ctx) error {

		return c.Status(200).JSON("pong")

	})

	// Mount user routes
	api.CreateUserRoutes(app)

	log.Fatal(app.Listen(":8080"))

}

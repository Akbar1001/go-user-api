package main

import (
	"context"
	"log"

	"go-user-api/config"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.LoadConfig()

	conn, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer conn.Close(context.Background())

	log.Println("Database Connected Successfully")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "User API Running",
		})
	})

	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
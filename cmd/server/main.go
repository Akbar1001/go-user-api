package main

import (
	
	"log"

	"go-user-api/config"

	sqlc "go-user-api/db/sqlc/generated"

	"go-user-api/internal/handler"
	"go-user-api/internal/repository"
	"go-user-api/internal/routes"
	"go-user-api/internal/service"
	"go-user-api/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go-user-api/internal/middleware"
)

func main() {

	err := logger.InitLogger()

	if err != nil {
		log.Fatal(err)
	}

	defer logger.Log.Sync()

	cfg := config.LoadConfig()

	conn, err := config.ConnectDB(cfg)

	if err != nil {
		log.Fatal("database connection failed:", err)
	}

	defer conn.Close()

	log.Println("Database Connected Successfully")

	queries := sqlc.New(conn)

	userRepo := repository.NewUserRepository(
		queries,
	)

	userService := service.NewUserService(
		userRepo,
	)

	userHandler := handler.NewUserHandler(
		userService,
	)

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	app.Get("/", func(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "API Running",
	})
})

	routes.SetupUserRoutes(
		app,
		userHandler,
	)

	log.Fatal(
		app.Listen(":" + cfg.ServerPort),
	)
}
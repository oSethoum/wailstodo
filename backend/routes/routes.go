package routes

import (
	"app/backend/handlers"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/websocket/v2"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	handlers.New()

	api := app.Group("/api")

	api.Get("/subscribe", websocket.New(handlers.Subscribe))

	app.Get("/docs/*", swagger.HandlerDefault)

	api.Get("/todos/", handlers.QueryTodo)
	api.Get("/todos/:id", handlers.OneTodo)
	api.Post("/todos/", handlers.CreateTodo)
	api.Patch("/todos/:id", handlers.UpdateTodo)
	api.Delete("/todos/:id", handlers.DeleteTodo)

	return app
}

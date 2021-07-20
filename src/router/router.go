package router

import (
	"github.com/derektruong/distribute-article-service/src/handler"
	"github.com/derektruong/distribute-article-service/src/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/account")

	user.Get("/:id", handler.GetAccount)
	user.Post("/", handler.CreateAccount)
	user.Patch("/:id", middleware.Protected(), handler.UpdateAccount)
	user.Delete("/:id", middleware.Protected(), handler.DeleteAccount)

	// Post
	product := api.Group("/product")
	product.Get("/", handler.GetAllPosts)
	product.Get("/:id", handler.GetPost)
	product.Post("/:id", middleware.Protected(), handler.CreatePost)
	product.Delete("/:id", middleware.Protected(), handler.DeletePost)
}

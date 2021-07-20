package router

import (
	"net/http"

	"github.com/derektruong/distribute-article-service/src/handler"
	"github.com/derektruong/distribute-article-service/src/handler/account"
	"github.com/derektruong/distribute-article-service/src/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App, cl *http.Client) {
	// Index page
	app.Get("/", handler.QuotesHandler(cl))

	// Render sign in page and welcome
	acc := app.Group("/account")
	acc.Get("/", account.RenderSignInHandler)
	acc.Get("/welcome", account.RenderWelcomeHandler)

	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)
	api.Get("/checkloggedin", handler.IsLoggedIn)

	// Sign up - sign in API
	api.Post("/signup", account.SignUpHandler)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// Account
	account := api.Group("/account")

	account.Get("/:id", handler.GetAccount)
	account.Post("/", handler.CreateAccount)
	account.Patch("/:id", middleware.Protected(), handler.UpdateAccount)
	account.Delete("/:id", middleware.Protected(), handler.DeleteAccount)

	// Post
	product := api.Group("/product")
	product.Get("/", handler.GetAllPosts)
	product.Get("/:id", handler.GetPost)
	product.Post("/:id", middleware.Protected(), handler.CreatePost)
	product.Delete("/:id", middleware.Protected(), handler.DeletePost)
}

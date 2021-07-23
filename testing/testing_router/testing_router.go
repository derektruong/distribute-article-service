package testing_router

import (

	"github.com/derektruong/distribute-article-service/testing/testing_handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupTestingRoutes(app *fiber.App) {


	app.Get("/editor", testing_handler.EditorHandler)
	
}
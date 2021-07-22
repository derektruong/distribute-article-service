package router

import (
	"net/http"

	"github.com/derektruong/distribute-article-service/src/config"
	"github.com/derektruong/distribute-article-service/src/handler/news"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupNewsRoutes(app *fiber.App, cl *http.Client) {

	// Render news page
	newsapi := news.NewClient(cl, config.Config("NEWS_API_KEY"), 12)

	app.Get("/search", news.SearchHandler(newsapi))
	app.Get("/headlines", news.HeadLinesHandler(newsapi))
	app.Get("/stocks", news.StocksHandler(newsapi))
	app.Get("/technology", news.TechHandler(newsapi))
	app.Get("/science", news.ScienceHandler(newsapi))
	app.Get("/sport", news.SportHandler(newsapi))
	
}
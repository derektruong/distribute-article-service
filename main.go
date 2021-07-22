package main

import (
	"log"
	"net/http"
	"time"

	"github.com/derektruong/distribute-article-service/src/database"
	"github.com/derektruong/distribute-article-service/src/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	_ "gorm.io/driver/mysql"
)

func main() {
	// Init client connection
	myClient := &http.Client{Timeout: 10 * time.Second}

	// Init routers with fiber
	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static files
	app.Static("/", "./public")
	
	app.Use(cors.New())

	// Connect to database
	database.ConnectDB()

	router.SetupRoutes(app, myClient)
	router.SetupNewsRoutes(app, myClient)
	
	// handle custom 404 responses
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("general/notfound", nil)
	})
	log.Fatal(app.Listen(":3000"))

	defer func ()  {
		sqlDB, err := database.DB.DB()
		if err != nil {
			log.Fatal(err)
		} else {
			sqlDB.Close()
		}
	}()
}
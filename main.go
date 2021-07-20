package main

import (
	"github.com/derektruong/distribute-article-service/src/database"
	"github.com/derektruong/distribute-article-service/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "gorm.io/driver/mysql"
)

func main() {
	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")
	
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
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
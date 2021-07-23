package testing_handler

import (

	"github.com/gofiber/fiber/v2"
)

func EditorHandler(c *fiber.Ctx) error {
	return c.Render("testing/index", nil)
}
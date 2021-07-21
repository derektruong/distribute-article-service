package account

import (
	// "net/mail"

	"github.com/derektruong/distribute-article-service/src/database"
	"github.com/derektruong/distribute-article-service/src/model"

	"github.com/gofiber/fiber/v2"
)

func RenderSignInHandler(c *fiber.Ctx) error {
	return c.Render("account/sign_in_up", nil)
}

func RenderWelcomeHandler(c *fiber.Ctx) error {
	return c.Render("account/welcome", fiber.Map{
		"Name": c.Locals("nameRegistered"),
	})
}

func SignUpHandler(c *fiber.Ctx) error {

	db := database.DB

	acc := &model.Account{
		IDRole: 3,
	}

	if err := c.BodyParser(acc); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	if r := db.First(&acc, "email = ?", acc.Email).Row(); r != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Email already exists", "data": nil})
	}

	// if _, isEmail := mail.ParseAddress(acc.Email); isEmail != nil {
	//     return c.Status(500).JSON(fiber.Map{"status": "error", "message": "This email is invalid", "data": nil})
	// }

	hash, err := hashPassword(acc.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	acc.Password = hash
	if err := db.Create(&acc).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create account", "data": err})
	}

	c.Locals("nameRegistered", acc.Name)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "created account", "data": nil})
}

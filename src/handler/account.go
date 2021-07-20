package handler

import (
	"github.com/derektruong/distribute-article-service/src/database"
	"github.com/derektruong/distribute-article-service/src/model"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func validAccount(id string, p string) bool {
	db := database.DB
	var acc model.Account
	db.First(&acc, id)
	if acc.Name == "" {
		return false
	}
	if !CheckPasswordHash(p, acc.Password) {
		return false
	}
	return true
}

// GetAccount get a user
func GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.Account
	db.Find(&user, id)
	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})
}

// CreateAccount new user
func CreateAccount(c *fiber.Ctx) error {
	type NewAccount struct {
		Name string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(model.Account)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newAccount := NewAccount{
		Email:    user.Email,
		Name: user.Name,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newAccount})
}

// UpdateAccount update user
func UpdateAccount(c *fiber.Ctx) error {
	type UpdateAccountInput struct {
		Name string `json:"name"`
	}
	var uui UpdateAccountInput
	if err := c.BodyParser(&uui); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})
	}

	db := database.DB
	var user model.Account

	db.First(&user, id)
	user.Name = uui.Name
	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "Account successfully updated", "data": user})
}

// DeleteAccount delete user
func DeleteAccount(c *fiber.Ctx) error {
	c.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjY5Nzg4MjIsIm5hbWUiOiJkdWMiLCJ1c2VyX2lkIjozN30.zy8YIfijWVaLzuZRCCTQRa5wk_emX9KPAmKS1bNG40U")
	
	type PasswordInput struct {
		Password string `json:"password"`
	}
	var pi PasswordInput
	if err := c.BodyParser(&pi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")
	token := c.Locals("user").(*jwt.Token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid token id", "data": nil})

	}

	if !validAccount(id, pi.Password) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})

	}

	db := database.DB
	var user model.Account

	db.First(&user, id)

	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "Account successfully deleted", "data": nil})
}

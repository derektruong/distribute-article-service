package account

import (
	"errors"
	"strings"
	"time"

	"github.com/derektruong/distribute-article-service/src/config"
	"github.com/derektruong/distribute-article-service/src/database"
	"github.com/derektruong/distribute-article-service/src/model"

	"github.com/golang-jwt/jwt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*model.Account, error) {
	db := database.DB
	var user model.Account
	if err := db.Where(&model.Account{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func IsLoggedIn(c *fiber.Ctx) error {

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	nameUser := string(claims["name"].(string))
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "logged in",
		"name_user": nameUser,
	})
}

// Login get user and password
func Login(c *fiber.Ctx) error {
	db := database.DB

	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type AccountData struct {
		ID       uint   `json:"id"`
		Role     string `json:"role"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	var ad AccountData
	var role model.RoleAccount

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	identity := input.Email
	pass := input.Password

	email, err := getUserByEmail(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	if email == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Account not found", "data": err})
	} else {
		if err := db.First(&role, "id = ?", email.IDRole).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Account not found", "data": err})
		}

		ad = AccountData{
			ID:       email.ID,
			Role:     role.RoleName,
			Name:     email.Name,
			Email:    email.Email,
			Password: email.Password,
		}
	}

	if !CheckPasswordHash(pass, ad.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = ad.Name
	claims["role"] = ad.Role
	claims["user_id"] = ad.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else {
		// Init jwt token for cookies
		jwtToken := strings.Split(t, ".")
		headerCookie := &fiber.Cookie{
			Name:     config.Config("JWT_TOKEN_HEADER"),
			Value:    jwtToken[0],
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Hour * 72),
		}
		payloadCookie := &fiber.Cookie{
			Name:     config.Config("JWT_TOKEN_PAYLOAD"),
			Value:    jwtToken[1],
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Hour * 72),
		}
		secretCookie := &fiber.Cookie{
			Name:     config.Config("JWT_TOKEN_SECRET"),
			Value:    jwtToken[2],
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Hour * 72),
		}

		// Set Cookie
		c.Cookie(headerCookie)
		c.Cookie(payloadCookie)
		c.Cookie(secretCookie)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func Logout(c *fiber.Ctx) error {
	// Init jwt token for cookies
	headerCookie := &fiber.Cookie{
		Name:     config.Config("JWT_TOKEN_HEADER"),
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-(time.Hour * 72)),
	}
	payloadCookie := &fiber.Cookie{
		Name:     config.Config("JWT_TOKEN_PAYLOAD"),
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-(time.Hour * 72)),
	}
	secretCookie := &fiber.Cookie{
		Name:     config.Config("JWT_TOKEN_SECRET"),
		Value:    "",
		HTTPOnly: true,
		Expires:  time.Now().Add(-(time.Hour * 72)),
	}

	// Set Cookie
	c.Cookie(headerCookie)
	c.Cookie(payloadCookie)
	c.Cookie(secretCookie)

	return c.JSON(fiber.Map{"status": "success", "message": "logged out successfully", "data": nil})
}

func GetJWTToken(c *fiber.Ctx) error {
	headerCookie := c.Cookies(config.Config("JWT_TOKEN_HEADER"))
	payloadCookie := c.Cookies(config.Config("JWT_TOKEN_PAYLOAD"))
	secretCookie := c.Cookies(config.Config("JWT_TOKEN_SECRET"))

	if headerCookie == "" || payloadCookie == "" || secretCookie == "" {
		return c.JSON(fiber.Map{"status": "error", "message": "unauthorized", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "got jwt", "data": headerCookie + "." + payloadCookie + "." + secretCookie})

}

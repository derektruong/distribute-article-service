package handler

import (
	"github.com/derektruong/distribute-article-service/src/database"
	"github.com/derektruong/distribute-article-service/src/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllPosts query all posts
func GetAllPosts(c *fiber.Ctx) error {
	db := database.DB
	var posts []model.Post
	db.Find(&posts)
	return c.JSON(fiber.Map{"status": "success", "message": "All posts", "data": posts})
}

// GetPost query post
func GetPost(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var post model.Post
	db.Find(&post, id)
	if post.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No post found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Post found", "data": post})
}

// CreatePost new Post
func CreatePost(c *fiber.Ctx) error {
	db := database.DB
	post := &model.Post {
		
	}
	if err := c.BodyParser(post); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create post", "data": err})
	}
	db.Create(&post)
	return c.JSON(fiber.Map{"status": "success", "message": "Created post", "data": post})
}

// DeletePost delete post
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var post model.Post
	db.First(&post, id)
	if post.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No post found with ID", "data": nil})

	}
	db.Delete(&post)
	return c.JSON(fiber.Map{"status": "success", "message": "Post successfully deleted", "data": nil})
}
package middleware

import (
	"time"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rogerhendricks/go-vue/models"
	"github.com/rogerhendricks/go-vue/database"
)


func AuthRequired(c *fiber.Ctx) error {
    db := database.InitDB()
    token := c.Cookies("session_token")

    // Check if the token is present
    if token == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Unauthorized",
        })
    }

    // Query the database for the session
    var session models.Session
    if err := db.Where("token = ?", token).First(&session).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Unauthorized",
        })
    }

    // Check if the session is expired
    if session.ExpiresAt.Before(time.Now()) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "message": "Session expired",
        })
    }

    // Set user ID in locals for further use
    c.Locals("userID", session.UserID)

    return c.Next()
}

// func IsAdmin(db *gorm.DB) func(*fiber.Ctx) error {

// 	return func(c *fiber.Ctx) error {
// 		user := c.Locals("user").(*models.User)
// 		if user.Role != "admin" {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"message": "Unauthorized",
// 			})
// 		}
// 		return c.Next()
// 	}
// }


package controllers

import (
    "log"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
	"golang.org/x/crypto/bcrypt"
)

var store = session.New(session.Config{
    KeyLookup:  "cookie:session_id",
    CookieSecure:   true,
    CookieHTTPOnly: true,
    CookieSameSite: "Strict",
})


func RegisterUser(c *fiber.Ctx) error {
	db := database.InitDB()

    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Check if username already exists
    var existingUser models.User
    if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
        return c.Status(400).JSON(fiber.Map{"error": "Username already exists"})
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

    user.Password = string(hashedPassword)

    result := db.Create(user)
    if result.Error != nil {
        log.Printf("Error creating user: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not register user"})
    }

	return c.JSON(fiber.Map{"message": "User registered successfully", "user": user})
}

func LoginUser(c *fiber.Ctx) error {
    db := database.InitDB()

    loginData := new(struct {
        Username string `json:"username"`
        Password string `json:"password"`
    })
    if err := c.BodyParser(loginData); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }
    var user models.User
    result := db.Where("username = ?", loginData.Username).First(&user)
    if result.Error != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
    }
    log.Printf("Password comparison successful")
    sess, err := store.Get(c)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
    }
    sess.Set("username", user.Username)
    if err := sess.Save(); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
    }

    user.Password = "" // Don't send password back
    return c.JSON(fiber.Map{"message": "Login successful", "user": user})
}
func LogoutUser(c *fiber.Ctx) error {
    sess, err := store.Get(c)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
    }
    if err := sess.Destroy(); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
    }
    return c.JSON(fiber.Map{"message": "Logout successful"})
}

func UpdateUser(c *fiber.Ctx) error {
    db := database.InitDB()
    sess, err := store.Get(c)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
    }
    username := sess.Get("username")
    if username == nil {
        return c.Status(401).JSON(fiber.Map{"error": "Not logged in"})
    }

    user := new(models.User)
    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    var existingUser models.User
    result := db.Where("username = ?", username).First(&existingUser)
    if result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    // Only update the fields if they are not empty
    if user.Username != "" {
        existingUser.Username = user.Username
    }
    if user.Fullname != "" {
        existingUser.Fullname = user.Fullname
    }
    if user.Password != "" {
        existingUser.Password = user.Password
    }

    result = db.Save(&existingUser)
    if result.Error != nil {

        return c.Status(500).JSON(fiber.Map{"error": "Could not update user"})
    }

    existingUser.Password = "" // Don't send password back
    return c.JSON(fiber.Map{"message": "User updated successfully", "user": existingUser})
}

func GetUser(c *fiber.Ctx) error {
    db := database.InitDB()
    sess, err := store.Get(c)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
    }
    username := sess.Get("username")
    if username == nil {
        return c.Status(401).JSON(fiber.Map{"error": "Not logged in"})
    }

    var user models.User
    result := db.Where("username = ?", username).First(&user)
    if result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    user.Password = "" // Don't send password back
    return c.JSON(fiber.Map{"user": user})
}
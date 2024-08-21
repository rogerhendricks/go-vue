package main

import (
    "log"
	"time"
	"strings"
	"embed"
	"net/http"
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	// "github.com/gofiber/fiber/v2/middleware/session"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/csrf"
    "github.com/gofiber/fiber/v2/utils"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "github.com/rogerhendricks/go-vue/controllers"
    "github.com/rogerhendricks/go-vue/models"
)

//go:embed frontend/dist
var distDir embed.FS

// type User struct {
//     gorm.Model
//     // ID        uint      `gorm:"primarykey" json:"id"`
//     // CreatedAt time.Time `json:"created_at"`
//     // UpdatedAt time.Time `json:"updated_at"`
//     // DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
//     Username  string    `gorm:"uniqueIndex" json:"username"`
//     Fullname  string    `json:"fullname"`
//     Password  string    `json:"-"`
// }

var db *gorm.DB
// var store = session.New(session.Config{
//     KeyLookup:  "cookie:session_id",
//     CookieSecure:   true,
//     CookieHTTPOnly: true,
//     CookieSameSite: "Strict",
// })

func main() {
    var err error
    db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&models.User{})
    // testPassword := "testpassword"
    // hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(testPassword), bcrypt.DefaultCost)
    // log.Printf("Test hashed password: %s", string(hashedPassword))

    // err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(testPassword))
    // if err != nil {
    //     log.Printf("Test password comparison failed: %v", err)
    // } else {
    //     log.Printf("Test password comparison successful")
    // }
    
    app := fiber.New()
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:8000",
        AllowMethods:     "GET,POST,PUT,DELETE",
        AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
        AllowCredentials: true,
    }))
	app.Use(csrf.New(csrf.Config{
        KeyLookup:      "header:X-CSRF-Token",
        CookieName:     "csrf_",
        CookieSameSite: "Strict",
        Expiration:     1 * time.Hour,
        KeyGenerator:   utils.UUID,
    }))
	api := app.Group("/api")
    api.Post("/register", controllers.RegisterUser)
    api.Post("/login", controllers.LoginUser)
    api.Post("/logout", controllers.LogoutUser)
    api.Get("/user", controllers.GetUser)


    api.Get("/devices", controllers.GetDevices)
    api.Post("/devices", controllers.CreateDevice)
    api.Get("/devices/:id", controllers.GetDevice)
    api.Put("/devices/:id", controllers.UpdateDevice)
    api.Delete("/devices/:id", controllers.DeleteDevice)
    
	app.Use("/", filesystem.New(filesystem.Config{
		Root:   http.FS(distDir),
		PathPrefix: "frontend/dist",
		Index:  "index.html",
		Browse: false,
	}))

    app.Use(func(c *fiber.Ctx) error {
        if strings.HasPrefix(c.Path(), "/assets") {
            return c.Next()
        }
        return c.SendFile("frontend/dist/index.html")
    })
    // app.Get("/api/hello", func (c *fiber.Ctx) error {
    //     return c.SendString("Hello, World!")
    // })

    log.Fatal(app.Listen(":8000"))
}


// func registerUser(c *fiber.Ctx) error {
//     user := new(User)
//     if err := c.BodyParser(user); err != nil {
//         log.Printf("Error parsing registration data: %v", err)
//         return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
//     }

//     log.Printf("Registering user with original password: %s", user.Password)

//     // Check if username already exists
//     var existingUser User
//     if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
//         return c.Status(400).JSON(fiber.Map{"error": "Username already exists"})
//     }

//     hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//     if err != nil {
//         log.Printf("Error hashing password: %v", err)
//         return c.Status(500).JSON(fiber.Map{"error": "Could not hash password"})
//     }
//     user.Password = string(hashedPassword)

//     log.Printf("Registering user with hashed password: %s", user.Password)

//     result := db.Create(&user)
//     if result.Error != nil {
//         log.Printf("Error creating user: %v", result.Error)
//         return c.Status(500).JSON(fiber.Map{"error": "Could not register user"})
//     }
//     log.Printf("User registered successfully: %s", user.Username)
//     user.Password = "" // Don't send password back
//     return c.JSON(fiber.Map{"message": "User registered successfully", "user": user})
// }
// func loginUser(c *fiber.Ctx) error {
//     loginData := new(struct {
//         Username string `json:"username"`
//         Password string `json:"password"`
//     })
//     if err := c.BodyParser(loginData); err != nil {
//         log.Printf("Error parsing login data: %v", err)
//         return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
//     }
//     log.Printf("Login attempt for user: %s with password: %s", loginData.Username, loginData.Password)
//     var user User
//     result := db.Where("username = ?", loginData.Username).First(&user)
//     if result.Error != nil {
//         log.Printf("User not found: %s", loginData.Username)
//         return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
//     }

//     log.Printf("Stored hashed password: %s", user.Password)
//     log.Printf("Provided password: %s", loginData.Password)

//     err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
//     if err != nil {
//         log.Printf("Password comparison failed: %v", err)
//         return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
//     }
//     log.Printf("Password comparison successful")
//     sess, err := store.Get(c)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
//     }
//     sess.Set("username", user.Username)
//     if err := sess.Save(); err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
//     }

//     user.Password = "" // Don't send password back
//     return c.JSON(fiber.Map{"message": "Login successful", "user": user})
// }

// func logoutUser(c *fiber.Ctx) error {
//     sess, err := store.Get(c)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
//     }
//     if err := sess.Destroy(); err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
//     }
//     return c.JSON(fiber.Map{"message": "Logout successful"})
// }

// func getUser(c *fiber.Ctx) error {
//     sess, err := store.Get(c)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
//     }
//     username := sess.Get("username")
//     if username == nil {
//         return c.Status(401).JSON(fiber.Map{"error": "Not logged in"})
//     }

//     var user User
//     result := db.Where("username = ?", username).First(&user)
//     if result.Error != nil {
//         return c.Status(404).JSON(fiber.Map{"error": "User not found"})
//     }

//     user.Password = "" // Don't send password back
//     return c.JSON(fiber.Map{"user": user})
// }
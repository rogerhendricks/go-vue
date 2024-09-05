package main

import (
    "log"
	"time"
	"strings"
	"embed"
	"net/http"
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
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
var db *gorm.DB

func main() {
    var err error
    db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&models.User{}, &models.Device{}, &models.Lead{}, &models.Doctor{}, &models.Address{} ,&models.Patient{})

    
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

    SetupRoutes(app)

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

    log.Fatal(app.Listen(":8000"))
}

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
    // users
    api.Post("/register", controllers.RegisterUser)
    api.Post("/login", controllers.LoginUser)
    api.Post("/logout", controllers.LogoutUser)
    api.Get("/user", controllers.GetUser)
    api.Put("/user", controllers.UpdateUser)
	// devices
	api.Get("/devices", controllers.GetDevices)
	api.Post("/devices", controllers.CreateDevice)
	api.Get("/devices/:id", controllers.GetDevice)
	api.Put("/devices/:id", controllers.UpdateDevice)
	api.Delete("/devices/:id", controllers.DeleteDevice)
	// leads
	api.Get("/leads", controllers.GetLeads)
	api.Post("/leads", controllers.CreateLead)
	api.Get("/leads/:id", controllers.GetLead)
	api.Put("leads/:id", controllers.UpdateLead)
	api.Delete("leads/:id", controllers.DeleteLead)
	// Doctors
	api.Get("/doctors", controllers.GetDoctors)
	api.Post("/doctors", controllers.CreateDoctor)
	api.Get("/doctors/:id", controllers.GetDoctor)
	api.Put("/doctors/:id", controllers.UpdateDoctor)
	api.Delete("/doctors/:id", controllers.DeleteDoctor)
    // Patients
    api.Get("/patients", controllers.GetPatients)
    app.Get("/api/patients/search", controllers.SearchPatients)
    api.Post("/patients", controllers.CreatePatient)
    api.Get("/patients/:id", controllers.GetPatient)
    api.Put("/patients/:id", controllers.UpdatePatient)
    api.Delete("/patients/:id", controllers.DeletePatient)
    // Lists
    api.Get("/deviceList", controllers.GetDevicesList)
    api.Get("/doctorsList", controllers.GetDoctorsList)
}
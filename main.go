package main

import (
	"embed"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/rogerhendricks/go-vue/controllers"
	"github.com/rogerhendricks/go-vue/middleware"
	"github.com/rogerhendricks/go-vue/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "github.com/rogerhendricks/go-vue/sessionstore"
)

//go:embed frontend/dist
var distDir embed.FS
var db *gorm.DB


var store = session.New(session.Config{
    Expiration: 24 * time.Hour, // Set session expiration time
})


func main() {
    var err error
    db, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database")
    }

    // Migrate the schema
    err = db.AutoMigrate(
        &models.User{}, &models.Device{},
        &models.Lead{}, &models.Doctor{},
        &models.Address{} ,&models.Patient{},
        &models.ImplantedDevice{}, &models.ImplantedLead{},
        &models.Report{}, &models.Session{})
    if err != nil {
            log.Fatalf("Failed to migrate Report: %v", err)
    }

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
        // Session:           session.Store,
        SessionKey:        "fiber.csrf.token",
        HandlerContextKey: "fiber.csrf.handler",
    }))


    SetupRoutes(app)

    app.Static("/uploads", "./uploads")

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
    api := app.Group("/api", middleware.AuthRequired())
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
    api.Get("/search/patients", controllers.SearchPatients)
    api.Post("/patients", controllers.CreatePatient)
    api.Get("/patients/:id", controllers.GetPatient)
    api.Put("/patients/:id", controllers.UpdatePatient)
    api.Delete("/patients/:id", controllers.DeletePatient)
    // Implanted Devices
    api.Get("/:id/implantedDevices", controllers.GetImplantedDevices)
    api.Post("/:id/implantedDevices", controllers.CreateImplantedDevice)
    api.Get("/implantedDevices/:id", controllers.GetImplantedDevice)
    api.Put("/implantedDevices/:id", controllers.UpdateImplantedDevice)
    api.Delete("/implantedDevices/:id", controllers.DeleteImplantedDevice)
    // Implanted Leads
    api.Get("/:id/implantedLeads", controllers.GetImplantedLeads)
    api.Post("/:id/implantedLeads", controllers.CreateImplantedLead)
    api.Get("/implantedLeads/:id", controllers.GetImplantedLead)
    api.Put("/implantedLeads/:id", controllers.UpdateImplantedLead)
    api.Delete("/implantedLeads/:id", controllers.DeleteImplantedLead)

    // Reports
    api.Get("/:id/reports", controllers.GetReports)
    api.Post("/:id/reports", controllers.CreateReport)
    api.Get("/reports/:id", controllers.GetReport)
    api.Put("/reports/:id", controllers.UpdateReport)
    api.Delete("/reports/:id", controllers.DeleteReport)


    // Lists
    api.Get("/deviceList", controllers.GetDevicesList)
    api.Get("/leadList", controllers.GetLeadsList)
    api.Get("/doctorsList", controllers.GetDoctorsList)
}

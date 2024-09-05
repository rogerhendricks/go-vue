// package main

// import (
//     "github.com/gofiber/fiber/v2"
//     "github.com/rogerhendricks/go-vue/controllers"
// )

// func SetupRoutes(app *fiber.App) {
//     api := app.Group("/api")
//     // users
//     api.Post("/register", controllers.RegisterUser)
//     api.Post("/login", controllers.LoginUser)
//     api.Post("/logout", controllers.LogoutUser)
//     api.Get("/user", controllers.GetUser)
//     api.Put("/user", controllers.UpdateUser)
// 	// devices
// 	api.Get("/devices", controllers.GetDevices)
// 	api.Post("/devices", controllers.CreateDevice)
// 	api.Get("/devices/:id", controllers.GetDevice)
// 	api.Put("/devices/:id", controllers.UpdateDevice)
// 	api.Delete("/devices/:id", controllers.DeleteDevice)
// 	// leads
// 	api.Get("/leads", controllers.GetLeads)
// 	api.Post("/leads", controllers.CreateLead)
// 	api.Get("/leads/:id", controllers.GetLead)
// 	api.Put("leads/:id", controllers.UpdateLead)
// 	api.Delete("leads/:id", controllers.DeleteLead)
// 	// Doctors
// 	api.Get("/doctors", controllers.GetDoctors)
// 	api.Post("/doctors", controllers.CreateDoctor)
// 	api.Get("/doctors/:id", controllers.GetDoctor)
// 	api.Put("/doctors/:id", controllers.UpdateDoctor)
// 	api.Delete("/doctors/:id", controllers.DeleteDoctor)
// }

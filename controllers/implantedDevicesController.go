package controllers

import (
	"log"
    "github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
)

func GetImplantedDevices(c *fiber.Ctx) error {
	db := database.InitDB()

	patientId := c.Params("id")
	var implantedDevices []models.ImplantedDevice
	result := db.Preload("Device").Preload("Doctor").Where("patient_id = ?", patientId).Order("implant_date desc").Find(&implantedDevices)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted devices"})
	}
	return c.JSON(fiber.Map{"implantedDevices": implantedDevices})
}

func GetImplantedDevice(c *fiber.Ctx) error {
	db := database.InitDB()

	id := c.Params("id")
	var implantedDevice models.ImplantedDevice
	result := db.Where("id = ?", id).First(&implantedDevice)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted device"})
	}
	return c.JSON(fiber.Map{"implantedDevice": implantedDevice})
}	

func CreateImplantedDevice(c *fiber.Ctx) error {
	db := database.InitDB()

	implantedDevice := new(models.ImplantedDevice)
	if err := c.BodyParser(implantedDevice); err != nil {
		log.Printf("Error parsing implanted device data: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	result := db.Create(implantedDevice)
	if result.Error != nil {
		log.Printf("Error creating implanted device: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not create implanted device"})
	}
	return c.JSON(fiber.Map{"implantedDevice": implantedDevice})
}

func UpdateImplantedDevice(c *fiber.Ctx) error {
	db := database.InitDB()

	id := c.Params("id")
	existingImplantedDevice := new(models.ImplantedDevice)
	result := db.Where("id = ?", id).First(&existingImplantedDevice)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted device"})
	}
	updatedImplantedDevice := new(models.ImplantedDevice)
	if err := c.BodyParser(updatedImplantedDevice); err != nil {
		log.Printf("Error parsing implanted device data: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	result = db.Model(&existingImplantedDevice).Updates(updatedImplantedDevice)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update implanted device"})
	}
	result = db.Preload("Device").
	Preload("Doctor").
	Preload("Patient").
	Where("id = ?", id).
	First(&existingImplantedDevice)

if result.Error != nil {
	return c.Status(500).JSON(fiber.Map{"error": "Could not fetch updated implanted device with associations"})
}
	return c.JSON(fiber.Map{"implantedDevice": existingImplantedDevice})
}

func DeleteImplantedDevice(c *fiber.Ctx) error {
	db := database.InitDB()

	id := c.Params("id")
	var implantedDevice models.ImplantedDevice
	result := db.Where("id = ?", id).First(&implantedDevice)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted device"})
	}
	result = db.Delete(&implantedDevice)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete implanted device"})
	}
	return c.JSON(fiber.Map{"message": "Implanted device deleted successfully"})
}
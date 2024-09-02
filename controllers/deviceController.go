package controllers

import (
    "log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
)

func GetDevices(c *fiber.Ctx) error {
    db := database.InitDB()

    var devices []models.Device
    result := db.Find(&devices)
    if result.Error != nil {
        log.Printf("Error fetching devices: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch devices"})
    }

    return c.JSON(fiber.Map{"devices": devices})
}

func GetDevicesList(c *fiber.Ctx) error {
    db := database.InitDB()

    var devices []models.Device
    result := db.Select("ID", "manufacturer", "name", "type").Find(&devices)
    if result.Error != nil {
        log.Printf("Error fetching devices: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch list of devices"})
    }

    return c.JSON(fiber.Map{"devices": devices})
}

func GetDevice(c *fiber.Ctx) error {
    db := database.InitDB()

    deviceID := c.Params("id")
    var device models.Device
    result := db.Where("id = ?", deviceID).First(&device)
    if result.Error != nil {
        log.Printf("Error fetching device: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch device"})
    }

    return c.JSON(fiber.Map{"device": device})
}

func CreateDevice(c *fiber.Ctx) error {
    db := database.InitDB()

    device := new(models.Device)
    if err := c.BodyParser(device); err != nil {
        log.Printf("Error parsing device data: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }
    log.Printf("Device data: %v", device)
    result := db.Create(device)
    if result.Error != nil {
        log.Printf("Error creating device: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not create device"})
    }

    return c.JSON(fiber.Map{"message": "Device created successfully", "device": device})
}

func UpdateDevice(c *fiber.Ctx) error {
    db := database.InitDB()

    deviceID := c.Params("id")

    existingDevice := new(models.Device)
    result := db.Where("id = ?", deviceID).First(existingDevice)
    if result.Error != nil {
        log.Printf("Error fetching device: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch device"})
    }

    // Parse request body into a new device object
    updatedDevice := new(models.Device)
    if err := c.BodyParser(updatedDevice); err != nil {
        log.Printf("Error parsing device data: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Update fields of the existing device with the parsed data
    existingDevice.Manufacturer = updatedDevice.Manufacturer
    existingDevice.DeviceModel = updatedDevice.DeviceModel
    existingDevice.Name = updatedDevice.Name
    existingDevice.Type = updatedDevice.Type
    existingDevice.Header = updatedDevice.Header
    existingDevice.HasHazard = updatedDevice.HasHazard
    existingDevice.IsMri = updatedDevice.IsMri
    existingDevice.Comment = updatedDevice.Comment
    // Continue updating other fields as necessary

    // Save the updated device to the database
    result = db.Save(existingDevice)
    if result.Error != nil {
        log.Printf("Error updating device: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not update device"})
    }

    return c.JSON(fiber.Map{"message": "Device updated successfully", "device": existingDevice})
}


func DeleteDevice(c *fiber.Ctx) error {
    db := database.InitDB()

    deviceID := c.Params("id")
    var device models.Device
    result := db.Where("id = ?", deviceID).First(&device)
    if result.Error != nil {
        log.Printf("Error fetching device: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch device"})
    }

    result = db.Delete(&device)
    if result.Error != nil {
        log.Printf("Error deleting device: %v", result.Error)		

		return c.Status(500).JSON(fiber.Map{"error": "Could not delete device"})
    }

    return c.JSON(fiber.Map{"message": "Device deleted successfully"})
}	
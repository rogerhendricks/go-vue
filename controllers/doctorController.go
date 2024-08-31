package controllers

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"

)

func GetDoctors(c *fiber.Ctx) error {
	db := database.InitDB()

	var doctors []models.Doctor
	result := db.Select("id", "Name", "Phone").Find(&doctors)
	if result.Error != nil {
		log.Printf("Error fetching doctors: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch doctors"})
	}

	return c.JSON(fiber.Map{"doctors": doctors})
}

func GetDoctor(c *fiber.Ctx) error {
	db := database.InitDB()

	doctorID := c.Params("id")
	var doctor models.Doctor
	result := db.Preload("Addresses").Where("id = ?", doctorID).First(&doctor)
	if result.Error != nil {
		log.Printf("Error fetching doctor: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch doctor"})
	}

	return c.JSON(fiber.Map{"doctor": doctor})
}

func CreateDoctor(c *fiber.Ctx) error {
	db := database.InitDB()

	doctor := new(models.Doctor)
	if err := c.BodyParser(doctor); err != nil {
		log.Printf("Error parsing doctor data: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	result := db.Create(doctor)
	if result.Error != nil {
		log.Printf("Error creating doctor: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not create doctor"})
	}

	return c.JSON(fiber.Map{"message": "Doctor created successfully", "doctor": doctor})
}

func UpdateDoctor(c *fiber.Ctx) error {
	
	db := database.InitDB()

	doctorID := c.Params("id")

	existingDoctor := new(models.Doctor)
	result := db.Preload("Addresses").Where("id = ?", doctorID).First(existingDoctor)
	if result.Error != nil {
		log.Printf("Error fetching doctor: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch doctor"})
	}

	newDoctor := new(models.Doctor)
	if err := c.BodyParser(newDoctor); err != nil {
		log.Printf("Error parsing doctor data: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
    // Update doctor's information
    existingDoctor.Name = newDoctor.Name
	existingDoctor.Phone = newDoctor.Phone
	existingDoctor.Email = newDoctor.Email

	// Handle addresses
	for _, address := range newDoctor.Addresses {
		existingAddress := new(models.Address)
		if err := db.Where("id = ?", address.ID).First(existingAddress).Error; err != nil {
			// If address doesn't exist, create a new one
			existingDoctor.Addresses = append(existingDoctor.Addresses, address)
		} else {
			// If address exists, update it
			db.Model(existingAddress).Updates(address)
		}
	}

		// Handle addresses to delete
		for _, address := range existingDoctor.Addresses {
			found := false
			for _, newAddress := range newDoctor.Addresses {
				if address.ID == newAddress.ID {
					found = true
					break
				}
			}
			if !found {
				db.Delete(&address)
			}
		}
		
	result = db.Save(existingDoctor)
	if result.Error != nil {
		log.Printf("Error updating doctor: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not update doctor"})
	}

	return c.JSON(fiber.Map{"message": "Doctor database updated successfully", "doctor": existingDoctor})
}

func DeleteDoctor(c *fiber.Ctx) error {
	db := database.InitDB()

	doctorID := c.Params("id")

	var doctor models.Doctor
	result := db.Preload("Addresses").Where("id = ?", doctorID).First(&doctor)
	if result.Error != nil {
		log.Printf("Error fetching doctor: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch doctor"})
	}

	result = db.Delete(&doctor)
	if result.Error != nil {
		log.Printf("Error deleting doctor: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete doctor"})
	}

	return c.JSON(fiber.Map{"message": "Doctor deleted successfully"})
}
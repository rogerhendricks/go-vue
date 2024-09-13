package controllers

import (
	"log"
    "github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
)

func GetImplantedLeads(c *fiber.Ctx) error {
	db := database.InitDB()

	patientId := c.Params("id")
	var implantedLeads []models.ImplantedLead
	result := db.Preload("Lead").Preload("Doctor").Where("patient_id = ?", patientId).Order("implant_date desc").Find(&implantedLeads)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted leads"})
	}
	return c.JSON(fiber.Map{"implantedLeads": implantedLeads})
}

func GetImplantedLead(c *fiber.Ctx) error {
	db := database.InitDB()

	id := c.Params("id")
	var implantedLead models.ImplantedLead
	result := db.Where("id = ?", id).First(&implantedLead)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted lead"})
	}
	return c.JSON(fiber.Map{"implantedLead": implantedLead})
}

func CreateImplantedLead(c *fiber.Ctx) error {
	db := database.InitDB()

	implantedLead := new(models.ImplantedLead)
	if err := c.BodyParser(implantedLead); err != nil {
		log.Printf("Error parsing implanted lead data: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	result := db.Create(implantedLead)
	if result.Error != nil {
		log.Printf("Error creating implanted lead: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not create implanted lead"})
	}
	return c.JSON(fiber.Map{"implantedLead": implantedLead})
}

func UpdateImplantedLead(c *fiber.Ctx) error {
	db := database.InitDB()

	id := c.Params("id")
	existingImplantedLead := new(models.ImplantedLead)
	result := db.Where("id = ?", id).First(&existingImplantedLead)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted lead"})
	}
	updatedImplantedLead := new(models.ImplantedLead)
	if err := c.BodyParser(updatedImplantedLead); err != nil {
		log.Printf("Error parsing implanted lead data: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	result = db.Model(&existingImplantedLead).Updates(updatedImplantedLead)
	if result.Error != nil {
		log.Printf("Error updating implanted lead: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not update implanted lead"})
	}

	result = db.Preload("Lead").
	Preload("Doctor").
	Preload("Patient").
	Where("id = ?", id).
	First(&existingImplantedLead)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch updated implanted lead with associations"})
	}
	return c.JSON(fiber.Map{"implantedLead": existingImplantedLead})
}

func DeleteImplantedLead(c *fiber.Ctx) error {
	db := database.InitDB()

	id := c.Params("id")
	var implantedLead models.ImplantedLead
	result := db.Where("id = ?", id).First(&implantedLead)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch implanted lead"})
	}
	result = db.Delete(&implantedLead)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete implanted lead"})
	}
	return c.SendStatus(204)
}
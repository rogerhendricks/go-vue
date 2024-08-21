package controllers
import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
)

func GetLeads(c *fiber.Ctx) error{
	db := database.InitDB()
	var leads []models.Lead
	result := db.Find(&leads)
    if result.Error != nil {
        log.Printf("Error fetching leads: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch leads"})
    }

    return c.JSON(fiber.Map{"leads": leads})
}

func GetLead(c *fiber.Ctx) error {
    db := database.InitDB()

    leadID := c.Params("id")
    var lead models.Lead
    result := db.Where("id = ?", leadID).First(&lead)
    if result.Error != nil {
        log.Printf("Error fetching lead: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch lead"})
    }

    return c.JSON(fiber.Map{"lead": lead})
}

func CreateLead(c *fiber.Ctx) error {
    db := database.InitDB()

    lead := new(models.Lead)
    if err := c.BodyParser(lead); err != nil {
        log.Printf("Error parsing lead data: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }
    log.Printf("lead data: %v", lead)
    result := db.Create(lead)
    if result.Error != nil {
        log.Printf("Error creating lead: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not create lead"})
    }

    return c.JSON(fiber.Map{"message": "Lead created successfully", "lead": lead})
}

func UpdateLead(c *fiber.Ctx) error {
    db := database.InitDB()

    leadID := c.Params("id")

    existingLead := new(models.Lead)
    result := db.Where("id = ?", leadID).First(existingLead)
    if result.Error != nil {
        log.Printf("Error fetching lead: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch lead"})
    }

    // Parse request body into a new lead object
    updatedLead := new(models.Lead)
    if err := c.BodyParser(updatedLead); err != nil {
        log.Printf("Error parsing lead data: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Update fields of the existing lead with the parsed data
    existingLead.Manufacturer = updatedLead.Manufacturer
    existingLead.LeadModel = updatedLead.LeadModel
    existingLead.Name = updatedLead.Name
    existingLead.Polarity = updatedLead.Polarity
    existingLead.HasHazard = updatedLead.HasHazard
    existingLead.IsMri = updatedLead.IsMri
    existingLead.Comment = updatedLead.Comment
    // Continue updating other fields as necessary

    // Save the updated lead to the database
    result = db.Save(existingLead)
    if result.Error != nil {
        log.Printf("Error updating lead: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not update lead"})
    }

    return c.JSON(fiber.Map{"message": "Lead updated successfully", "lead": existingLead})
}

func DeleteLead(c *fiber.Ctx) error {
    db := database.InitDB()

    leadID := c.Params("id")
    var lead models.Lead
    result := db.Where("id = ?", leadID).First(&lead)
    if result.Error != nil {
        log.Printf("Error fetching lead: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch lead"})
    }

    result = db.Delete(&lead)
    if result.Error != nil {
        log.Printf("Error deleting lead: %v", result.Error)		

		return c.Status(500).JSON(fiber.Map{"error": "Could not delete lead"})
    }

    return c.JSON(fiber.Map{"message": "Lead deleted successfully"})
}	
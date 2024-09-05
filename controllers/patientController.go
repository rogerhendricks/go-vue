package controllers

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
)

func GetPatients(c *fiber.Ctx) error {
	db := database.InitDB()

	var patients []models.Patient
	result := db.Select("id", "Name", "Phone").Find(&patients)
	if result.Error != nil {
		log.Printf("Error fetching patients: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch patients"})
	}

	return c.JSON(fiber.Map{"patients": patients})
}

func SearchPatients(c *fiber.Ctx) error {
    db := database.InitDB()


    searchTerm := c.Query("search", "")

    var patients []models.Patient
    db.Where("name LIKE ?", "%" + searchTerm + "%").Find(&patients)
    return c.JSON(patients)
}


func GetPatient(c *fiber.Ctx) error {
	db := database.InitDB()

	patientID := c.Params("id")
	var patient models.Patient
	result := db.Preload("Doctors").Where("id = ?", patientID).First(&patient)
	if result.Error != nil {
		log.Printf("Error fetching patient: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch patient"})
	}

	return c.JSON(fiber.Map{"patient": patient})
}

// func CreatePatient(c *fiber.Ctx) error {
// 	db := database.InitDB()
	
// 	patient := new(models.Patient)
// 	if err := c.BodyParser(patient); err != nil {
// 		log.Printf("Error parsing patient data: %v", err)
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
// 	}

// 	result := db.Create(patient)
// 	if result.Error != nil {
// 		log.Printf("Error creating patient: %v", result.Error)
// 		return c.Status(500).JSON(fiber.Map{"error": "Could not create patient"})
// 	}

// 	return c.JSON(fiber.Map{"message": "Patient created successfully", "patient": patient})
// }

func CreatePatient(c *fiber.Ctx) error {
    db := database.InitDB()
    
    patient := new(models.Patient)
    if err := c.BodyParser(patient); err != nil {
        log.Printf("Error parsing patient data: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    result := db.Create(patient)
    if result.Error != nil {
        log.Printf("Error creating patient: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not create patient"})
    }

    // Handle doctor associations
    for _, doctor := range patient.Doctors {
        existingDoctor := new(models.Doctor)
        if err := db.Where("id = ?", doctor.ID).First(existingDoctor).Error; err != nil {
            log.Printf("Error fetching doctor: %v", err)
        } else {
            db.Model(patient).Association("Doctors").Append(existingDoctor)
        }
    }

    return c.JSON(fiber.Map{"message": "Patient created successfully", "patient": patient})
}

func UpdatePatient(c *fiber.Ctx) error {
    db := database.InitDB()

    patientID := c.Params("id")

    patient := new(models.Patient)
    result := db.Preload("Doctors").Where("id = ?", patientID).First(&patient)
    if result.Error != nil {
        log.Printf("Error fetching patient: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch patient"})
    }

    newPatient := new(models.Patient)
    if err := c.BodyParser(&newPatient); err != nil {
		log.Println(c.Body())
        log.Printf("Error parsing patient data: %v", err)
        return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
    }

    patient.Name = newPatient.Name
    patient.Phone = newPatient.Phone
    patient.Email = newPatient.Email
    patient.Address = newPatient.Address
    patient.State = newPatient.State
    patient.City = newPatient.City
    patient.Zip = newPatient.Zip
    patient.Country = newPatient.Country
    patient.Sex = newPatient.Sex
    patient.Dob = newPatient.Dob
    patient.Active = newPatient.Active

    // Clear existing doctor relationships
    db.Model(&patient).Association("Doctors").Clear()

    // Add new doctor relationships
    for _, doctor := range newPatient.Doctors {
        existingDoctor := new(models.Doctor)
        if err := db.Where("id = ?", doctor.ID).First(existingDoctor).Error; err != nil {
            log.Printf("Error fetching doctor: %v", err)
        } else {
            db.Model(patient).Association("Doctors").Append(existingDoctor)
        }
    }

    result = db.Save(&patient)
    if result.Error != nil {
        log.Printf("Error updating patient: %v", result.Error)
        return c.Status(500).JSON(fiber.Map{"error": "Could not update patient"})
    }

    return c.JSON(fiber.Map{"message": "Patient updated successfully", "patient": patient})
}
// func UpdatePatient(c *fiber.Ctx) error {
	
// 	db := database.InitDB()

// 	patientID := c.Params("id")

// 	patient := new(models.Patient)
// 	result := db.Preload("doctors").Where("id = ?", patientID).First(&patient)
// 	if result.Error != nil {
// 		log.Printf("Error fetching patient: %v", result.Error)
// 		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch patient"})
// 	}

// 	newPatient := new(models.Patient)
// 	if err := c.BodyParser(&newPatient); err != nil {
// 		log.Printf("Error parsing patient data: %v", err)
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
// 	}

// 	patient.Name = newPatient.Name
// 	patient.Phone = newPatient.Phone
// 	patient.Email = newPatient.Email
// 	patient.Address = newPatient.Address
// 	patient.State = newPatient.State
// 	patient.City = newPatient.City
// 	patient.Zip = newPatient.Zip
// 	patient.Country = newPatient.Country
// 	patient.Sex = newPatient.Sex
// 	patient.Dob = newPatient.Dob
// 	patient.Active = newPatient.Active

//     // Clear existing doctor relationships
//     db.Model(&patient).Association("Doctors").Clear()

//     // Add new doctor relationships
//     for _, doctor := range newPatient.Doctors {
//         existingDoctor := new(models.Doctor)
//         if err := db.Where("id = ?", doctor.ID).First(existingDoctor).Error; err != nil {
//             log.Printf("Error fetching doctor: %v", err)
//         } else {
//             patient.Doctors = append(patient.Doctors, *existingDoctor)
//         }
//     }

// 	// for _, doctor := range newPatient.Doctors {
// 	// 	existingDoctor := new(models.Doctor)
// 	// 	if err := db.Where("id = ?", doctor.ID).First(patient).Error; err != nil {
// 	// 		log.Printf("Error fetching doctor: %v", err)
// 	// 		patient.Doctors = append(patient.Doctors, doctor)
// 	// 	} else {
// 	// 		db.Model(existingDoctor).Updates(doctor)
// 	// 	}
		
// 	// }

// 	result = db.Save(&patient)
// 	if result.Error != nil {
// 		log.Printf("Error updating patient: %v", result.Error)
// 		return c.Status(500).JSON(fiber.Map{"error": "Could not update patient"})
// 	}

// 	return c.JSON(fiber.Map{"message": "Patient updated successfully", "patient": patient})
// }

func DeletePatient(c *fiber.Ctx) error {
	db := database.InitDB()

	patientID := c.Params("id")
	var patient models.Patient
	result := db.Where("id = ?", patientID).First(&patient)
	if result.Error != nil {
		log.Printf("Error fetching patient: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch patient"})
	}

	// Delete doctor relationships
	db.Model(&patient).Association("Doctors").Clear()

	result = db.Delete(&patient)
	if result.Error != nil {
		log.Printf("Error deleting patient: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete patient"})
	}

	return c.JSON(fiber.Map{"message": "Patient deleted successfully"})
}


package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/rogerhendricks/go-vue/database"
	"github.com/rogerhendricks/go-vue/models"
)

func GetReports(c *fiber.Ctx) error {
	db := database.InitDB()

    patientId := c.Params("id")
	var reports []models.Report
	result := db.Select("ID", "report_date").Where("patient_id = ?", patientId).Order("report_date desc").Find(&reports)
	if result.Error != nil {
		log.Printf("Error fetching reports: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch reports"})
	}

	return c.JSON(fiber.Map{"reports": reports})
}

func GetReport(c *fiber.Ctx) error {
	db := database.InitDB()

	reportID := c.Params("id")
	var report models.Report
	result := db.Where("id = ?", reportID).First(&report)
	if result.Error != nil {
		log.Printf("Error fetching report: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch report"})
	}

	return c.JSON(fiber.Map{"report": report})
}

func CreateReport(c *fiber.Ctx) error {
    db := database.InitDB()

    // Get file from the form
    file, err := c.FormFile("file")
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "File is required")
    }

    // Get patient_id and report_date
    patientIDStr := c.FormValue("patient_id")
    reportDate := c.FormValue("report_date")
    
    // Convert patient_id to uint
    patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid patient ID")
    }

    // Create directories if they don't exist
    dirPath := filepath.Join("uploads", fmt.Sprint(patientID), reportDate)
    if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
    }

    // Generate the file path
    filePath := filepath.Join(dirPath, file.Filename)

    // Save the file
    if err := c.SaveFile(file, filePath); err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
    }

    // Save the report in the database
    report := models.Report{
        ReportDate: reportDate,
        PatientID:  uint(patientID),
        FilePath:   filePath,
    }
    if err := db.Create(&report).Error; err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to save report to database")
    }

    return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Report created successfully",
        "data":    report,
    })
}


func UpdateReport(c *fiber.Ctx) error {
	db := database.InitDB()

    // Get patient_id and report_date from the form
    patientIDStr := c.FormValue("patient_id")
    reportDate := c.FormValue("report_date")

    // Convert patient_id to uint
    patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid patient ID")
    }

    // Find the existing report
    var report models.Report
    if err := db.Where("patient_id = ? AND report_date = ?", patientID, reportDate).First(&report).Error; err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Report not found")
    }

    // Handle file upload (optional)
    file, err := c.FormFile("file")
    if err == nil {
        // If a new file is uploaded, remove the old file
        if report.FilePath != "" {
            os.Remove(report.FilePath) // Optionally handle errors here
        }

        // Create directories if they don't exist
        dirPath := filepath.Join("uploads", fmt.Sprint(patientID), reportDate)
        if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
        }

        // Generate new file path and save file
        newFilePath := filepath.Join(dirPath, file.Filename)
        if err := c.SaveFile(file, newFilePath); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
        }

        // Update file path in the report
        report.FilePath = newFilePath
    }

    // Optionally update other fields (e.g., report date)
    if newReportDate := c.FormValue("new_report_date"); newReportDate != "" {
        report.ReportDate = newReportDate
    }

    // Save the updated report
    if err := db.Save(&report).Error; err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to update report in database")
    }

    return c.JSON(fiber.Map{
        "status":  "success",
        "message": "Report updated successfully",
        "data":    report,
    })
}

func DeleteReport(c *fiber.Ctx) error {
	db := database.InitDB()

	reportID := c.Params("id")
	var report models.Report
	result := db.Where("id = ?", reportID).First(&report)
	if result.Error != nil {
		log.Printf("Error fetching report: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch report"})
	}

	result = db.Delete(&report)
	if result.Error != nil {
		log.Printf("Error deleting report: %v", result.Error)
		return c.Status(500).JSON(fiber.Map{"error": "Could not delete report"})
	}

	return c.JSON(fiber.Map{"message": "Report deleted successfully"})
}


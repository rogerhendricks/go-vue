package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
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
    file, _ := c.FormFile("file")
    hasFile := file != nil

    authorIdStr := c.FormValue("author_id")
    patientIDStr := c.FormValue("patient_id")
    reportDate := c.FormValue("report_date")
    reportType := c.FormValue("report_type")
    reportStatus := c.FormValue("report_status")
    currentDependency := c.FormValue("current_dependency")
    currentHeartRate := c.FormValue("current_heart_rate")
    
    // Convert patient_id to uint
    patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid patient ID")
    }

    // Convert author_id to uint
    authorId, err := strconv.ParseUint(authorIdStr, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid author ID")
    }
    var filePath string
    if hasFile {
        // Create directories if they don't exist
        dirPath := filepath.Join("uploads", fmt.Sprint(patientID), reportDate)
        if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
        }

        // Generate the file path
        filePath = filepath.Join(dirPath, file.Filename)

        // Save the file
        if err := c.SaveFile(file, filePath); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
        }
    }

    // // Create directories if they don't exist
    // dirPath := filepath.Join("uploads", fmt.Sprint(patientID), reportDate)
    // if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
    //     return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
    // }

    // // Generate the file path
    // filePath := filepath.Join(dirPath, file.Filename)

    // // Save the file
    // if err := c.SaveFile(file, filePath); err != nil {
    //     return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
    // }

    // Save the report in the database
    currentHeartRateInt, err := strconv.Atoi(currentHeartRate)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid current heart rate")
    }

    report := models.Report{
        AuthorID:          uint(authorId),
        ReportDate:        reportDate,
        ReportType:        reportType,
        ReportStatus:      reportStatus,
        CurrentHeartRate:  currentHeartRateInt,
        CurrentDependency: currentDependency,
        PatientID:         uint(patientID),
        FilePath:         filePath,
    }
    // if hasFile {
    //     report.FilePath = filePath
    // }
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

    reportId := c.Params("id")
    // Get patient_id and report_date from the form
    patientIDStr := c.FormValue("patient_id")
    reportDate := c.FormValue("report_date")

    // Convert patient_id to uint
    patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid patient ID")
    }

    // Find the existing report
    report := new(models.Report)
    if err := db.Where("id = ?", reportId).First(&report).Error; err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Report not found")
    }
    
    // Handle file upload
    file, err := c.FormFile("file")
    if err == nil {
        // If a new file is uploaded, choose to overwrite or append
        appendFile := c.FormValue("append") == "true"

        if !appendFile {
            // Overwrite: remove the old file
            if report.FilePath != "" {
                os.Remove(report.FilePath) // Optionally handle errors here
            }
        }

        // Create directories if they don't exist
        dirPath := filepath.Join("uploads", fmt.Sprint(patientID), reportDate)
        if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
        }

        newFilePath := filepath.Join(dirPath, file.Filename)

        // Save the file (either overwrite or append)
        if appendFile {
            // Append logic for PDF (or any file type)
            err = appendToFile(report.FilePath, file)
            if err != nil {
                return fiber.NewError(fiber.StatusInternalServerError, "Failed to append to file")
            }
        } else {
            // Overwrite the old file by saving the new one
            if err := c.SaveFile(file, newFilePath); err != nil {
                return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
            }
            report.FilePath = newFilePath
        }
    }

    // Optionally update other fields (e.g., report date)
    if reportDate := c.FormValue("report_date"); reportDate != "" {
        report.ReportDate = reportDate
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



// Function to append content to an existing file
func appendToFile(filePath string, fileHeader *multipart.FileHeader) error {
    // Open the existing file in append mode
    existingFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer existingFile.Close()

    // Open the uploaded file
    uploadedFile, err := fileHeader.Open()
    if err != nil {
        return err
    }
    defer uploadedFile.Close()

    // Read the contents of the uploaded file
    fileData := make([]byte, fileHeader.Size)
    _, err = uploadedFile.Read(fileData)
    if err != nil {
        return err
    }

    // Append the uploaded file's data to the existing file
    _, err = existingFile.Write(fileData)
    if err != nil {
        return err
    }

    return nil
}
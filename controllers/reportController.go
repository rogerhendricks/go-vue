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

// func CreateReport(c *fiber.Ctx) error {
// 	db := database.InitDB()

// 	// Get file from the form
// 	file, _ := c.FormFile("file")
// 	hasFile := file != nil

// 	authorIdStr := c.FormValue("author_id")
// 	patientIDStr := c.FormValue("patient_id")
// 	reportDate := c.FormValue("report_date")
// 	reportType := c.FormValue("report_type")
// 	reportStatus := c.FormValue("report_status")
// 	currentDependency := c.FormValue("current_dependency")
// 	currentHeartRate := c.FormValue("current_heart_rate")

// 	// Convert patient_id to uint
// 	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid patient ID")
// 	}

// 	// Convert author_id to uint
// 	authorId, err := strconv.ParseUint(authorIdStr, 10, 32)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid author ID")
// 	}
// 	var filePath string
// 	if hasFile {
// 		// Create directories if they don't exist
// 		dirPath := filepath.Join("uploads", fmt.Sprint(patientID), reportDate)
// 		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
// 			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create directory")
// 		}

// 		// Generate the file path
// 		filePath = filepath.Join(dirPath, file.Filename)

// 		// Save the file
// 		if err := c.SaveFile(file, filePath); err != nil {
// 			return fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
// 		}
// 	}
// 	// Save the report in the database
// 	currentHeartRateInt, err := strconv.Atoi(currentHeartRate)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid current heart rate")
// 	}

// 	report := models.Report{
// 		AuthorID:          uint(authorId),
// 		ReportDate:        reportDate,
// 		ReportType:        reportType,
// 		ReportStatus:      reportStatus,
// 		CurrentHeartRate:  currentHeartRateInt,
// 		CurrentDependency: currentDependency,
// 		PatientID:         uint(patientID),
// 		FilePath:          filePath,
// 	}

// 	if err := db.Create(&report).Error; err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "Failed to save report to database")
// 	}

//		return c.JSON(fiber.Map{
//			"status":  "success",
//			"message": "Report created successfully",
//			"data":    report,
//		})
//	}
func CreateReport(c *fiber.Ctx) error {
	db := database.InitDB()

	// Get file from the form
	file, _ := c.FormFile("file")
	hasFile := file != nil

	// Get form values
	authorIdStr := c.FormValue("author_id")
	patientIDStr := c.FormValue("patient_id")
	reportDate := c.FormValue("report_date")
	reportType := c.FormValue("report_type")
	reportStatus := c.FormValue("report_status")
	currentHeartRateStr := c.FormValue("current_heart_rate")
	currentRhythm := c.FormValue("current_rhythm")
	currentDependency := c.FormValue("current_dependency")
	mdc_idc__stat_ataf_burden_percentStr := c.FormValue("mdc_idc__stat_ataf_burden_percent")
	mdc_idc_set_brady_mode := c.FormValue("mdc_idc_set_brady_mode")
	mdc_idc_set_brady_lowrateStr := c.FormValue("mdc_idc_set_brady_lowrate")
	mdc_idc_set_brady_max_tracking_rateStr := c.FormValue("mdc_idc_set_brady_max_tracking_rate")
	mdc_idc_set_brady_max_sensor_rateStr := c.FormValue("mdc_idc_set_brady_max_sensor_rate")
	mdc_idc_dev_sav := c.FormValue("mdc_idc_dev_sav")
	mdc_idc_dev_pav := c.FormValue("mdc_idc_dev_pav")
	mdc_idc_stat_brady_ra_percent_pacedStr := c.FormValue("mdc_idc_stat_brady_ra_percent_paced")
	mdc_idc_stat_brady_rv_percent_pacedStr := c.FormValue("mdc_idc_stat_brady_rv_percent_paced")
	mdc_idc_stat_brady_lv_percent_pacedStr := c.FormValue("mdc_idc_stat_brady_lv_percent_paced")
	mdc_idc_stat_brady_biv_percent_pacedStr := c.FormValue("mdc_idc_stat_brady_biv_percent_paced")
	mdc_idc_batt_voltStr := c.FormValue("mdc_idc_batt_volt")
	mdc_idc_batt_remaining := c.FormValue("mdc_idc_batt_remaining")
	mdc_idc_batt_status := c.FormValue("mdc_idc_batt_status")
	mdc_idc_cap_charge_timeStr := c.FormValue("mdc_idc_cap_charge_time")
	mdc_idc_msmt_ra_impedance_meanStr := c.FormValue("mdc_idc_msmt_ra_impedance_mean")
	mdc_idc_msmt_ra_sensing_meanStr := c.FormValue("mdc_idc_msmt_ra_sensing_mean")
	mdc_idc_msmt_ra_thresholdStr := c.FormValue("mdc_idc_msmt_ra_threshold")
	mdc_idc_msmt_ra_pwStr := c.FormValue("mdc_idc_msmt_ra_pw")
	mdc_idc_msmt_rv_impedance_meanStr := c.FormValue("mdc_idc_msmt_rv_impedance_mean")
	mdc_idc_msmt_rv_sensing_meanStr := c.FormValue("mdc_idc_msmt_rv_sensing_mean")
	mdc_idc_msmt_rv_thresholdStr := c.FormValue("mdc_idc_msmt_rv_threshold")
	mdc_idc_msmt_rv_pwStr := c.FormValue("mdc_idc_msmt_rv_pw")
	mdc_idc_msmt_shock_impedanceStr := c.FormValue("mdc_idc_msmt_shock_impedance")
	mdc_idc_msmt_lv_impedance_meanStr := c.FormValue("mdc_idc_msmt_lv_impedance_mean")
	mdc_idc_msmt_lv_thresholdStr := c.FormValue("mdc_idc_msmt_lv_threshold")
	mdc_idc_msmt_lv_pwStr := c.FormValue("mdc_idc_msmt_lv_pw")
	comments := c.FormValue("comments")
	isCompletedStr := c.FormValue("is_completed")

	// Convert necessary form values to appropriate types
	authorId, err := strconv.ParseUint(authorIdStr, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid author ID")
	}
	patientID, err := strconv.ParseUint(patientIDStr, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid patient ID")
	}
	currentHeartRate, err := parseNullableInt(currentHeartRateStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid current heart rate")
	}
	mdc_idc__stat_ataf_burden_percent, err := parseNullableInt(mdc_idc__stat_ataf_burden_percentStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc__stat_ataf_burden_percent")
	}
	mdc_idc_set_brady_lowrate, err := parseNullableInt(mdc_idc_set_brady_lowrateStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_set_brady_lowrate")
	}
	mdc_idc_set_brady_max_tracking_rate, err := parseNullableInt(mdc_idc_set_brady_max_tracking_rateStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_set_brady_max_tracking_rate")
	}
	mdc_idc_set_brady_max_sensor_rate, err := parseNullableInt(mdc_idc_set_brady_max_sensor_rateStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_set_brady_max_sensor_rate")
	}
	mdc_idc_stat_brady_ra_percent_paced, err := parseNullableInt(mdc_idc_stat_brady_ra_percent_pacedStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_stat_brady_ra_percent_paced")
	}
	mdc_idc_stat_brady_rv_percent_paced, err := parseNullableInt(mdc_idc_stat_brady_rv_percent_pacedStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_stat_brady_rv_percent_paced")
	}
	mdc_idc_stat_brady_lv_percent_paced, err := parseNullableInt(mdc_idc_stat_brady_lv_percent_pacedStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_stat_brady_lv_percent_paced")
	}
	mdc_idc_stat_brady_biv_percent_paced, err := parseNullableInt(mdc_idc_stat_brady_biv_percent_pacedStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_stat_brady_biv_percent_paced")
	}
	mdc_idc_batt_volt, err := parseNullableFloat(mdc_idc_batt_voltStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_batt_volt")
	}
	// mdc_idc_batt_remaining, err := mdc_idc_batt_remainingStr
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_batt_remaining")
	// }
	mdc_idc_cap_charge_time, err := parseNullableFloat(mdc_idc_cap_charge_timeStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_cap_charge_time")
	}
	mdc_idc_msmt_ra_impedance_mean, err := parseNullableFloat(mdc_idc_msmt_ra_impedance_meanStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_ra_impedance_mean")
	}
	mdc_idc_msmt_ra_sensing_mean, err := parseNullableFloat(mdc_idc_msmt_ra_sensing_meanStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_ra_sensing_mean")
	}
	mdc_idc_msmt_ra_threshold, err := parseNullableFloat(mdc_idc_msmt_ra_thresholdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_ra_threshold")
	}
	mdc_idc_msmt_ra_pw, err := parseNullableFloat(mdc_idc_msmt_ra_pwStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_ra_pw")
	}
	mdc_idc_msmt_rv_impedance_mean, err := parseNullableFloat(mdc_idc_msmt_rv_impedance_meanStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_rv_impedance_mean")
	}
	mdc_idc_msmt_rv_sensing_mean, err := parseNullableFloat(mdc_idc_msmt_rv_sensing_meanStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_rv_sensing_mean")
	}
	mdc_idc_msmt_rv_threshold, err := parseNullableFloat(mdc_idc_msmt_rv_thresholdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_rv_threshold")
	}
	mdc_idc_msmt_rv_pw, err := parseNullableFloat(mdc_idc_msmt_rv_pwStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_rv_pw")
	}
	mdc_idc_msmt_shock_impedance, err := parseNullableFloat(mdc_idc_msmt_shock_impedanceStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_shock_impedance")
	}
	mdc_idc_msmt_lv_impedance_mean, err := parseNullableFloat(mdc_idc_msmt_lv_impedance_meanStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_lv_impedance_mean")
	}
	mdc_idc_msmt_lv_threshold, err := parseNullableFloat(mdc_idc_msmt_lv_thresholdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_lv_threshold")
	}
	mdc_idc_msmt_lv_pw, err := parseNullableFloat(mdc_idc_msmt_lv_pwStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid mdc_idc_msmt_lv_pw")
	}
	isCompleted, err := strconv.ParseBool(isCompletedStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid is_completed value")
	}

	// Debug statements to check the parsed values
	// fmt.Printf("mdc_idc_set_brady_max_tracking_rate: %v\n", *mdc_idc_set_brady_max_tracking_rate)


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

	// Create a new report instance
	report := models.Report{
		AuthorID:                             uint(authorId),
		ReportDate:                           reportDate,
		ReportType:                           reportType,
		ReportStatus:                         reportStatus,
		CurrentHeartRate:                     currentHeartRate,
		CurrentRhythm:                        currentRhythm,
		CurrentDependency:                    currentDependency,
		Mdc_idc__stat_ataf_burden_percent:    mdc_idc__stat_ataf_burden_percent,
		Mdc_idc_set_brady_mode:               mdc_idc_set_brady_mode,
		Mdc_idc_set_brady_lowrate:            mdc_idc_set_brady_lowrate,
		Mdc_idc_set_brady_max_tracking_rate:  mdc_idc_set_brady_max_tracking_rate,
		Mdc_idc_set_brady_max_sensor_rate:    mdc_idc_set_brady_max_sensor_rate,
		Mdc_idc_dev_sav:                      mdc_idc_dev_sav,
		Mdc_idc_dev_pav:                      mdc_idc_dev_pav,
		Mdc_idc_stat_brady_ra_percent_paced:  mdc_idc_stat_brady_ra_percent_paced,
		Mdc_idc_stat_brady_rv_percent_paced:  mdc_idc_stat_brady_rv_percent_paced,
		Mdc_idc_stat_brady_lv_percent_paced:  mdc_idc_stat_brady_lv_percent_paced,
		Mdc_idc_stat_brady_biv_percent_paced: mdc_idc_stat_brady_biv_percent_paced,
		Mdc_idc_batt_volt:                    mdc_idc_batt_volt,
		Mdc_idc_batt_remaining:               mdc_idc_batt_remaining,
		Mdc_idc_batt_status:                  mdc_idc_batt_status,
		Mdc_idc_cap_charge_time:              mdc_idc_cap_charge_time,
		Mdc_idc_msmt_ra_impedance_mean:       mdc_idc_msmt_ra_impedance_mean,
		Mdc_idc_msmt_ra_sensing_mean:         mdc_idc_msmt_ra_sensing_mean,
		Mdc_idc_msmt_ra_threshold:            mdc_idc_msmt_ra_threshold,
		Mdc_idc_msmt_ra_pw:                   mdc_idc_msmt_ra_pw,
		Mdc_idc_msmt_rv_impedance_mean:       mdc_idc_msmt_rv_impedance_mean,
		Mdc_idc_msmt_rv_sensing_mean:         mdc_idc_msmt_rv_sensing_mean,
		Mdc_idc_msmt_rv_threshold:            mdc_idc_msmt_rv_threshold,
		Mdc_idc_msmt_rv_pw:                   mdc_idc_msmt_rv_pw,
		Mdc_idc_msmt_shock_impedance:         mdc_idc_msmt_shock_impedance,
		Mdc_idc_msmt_lv_impedance_mean:       mdc_idc_msmt_lv_impedance_mean,
		Mdc_idc_msmt_lv_threshold:            mdc_idc_msmt_lv_threshold,
		Mdc_idc_msmt_lv_pw:                   mdc_idc_msmt_lv_pw,
		Comments:                             comments,
		IsCompleted:                          isCompleted,
		FilePath:                             filePath,
		PatientID:                            uint(patientID),
	}

	// Save the report to the database
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

// Helper function to parse nullable integers
func parseNullableInt(value string) (*int, error) {
    if value == "" || value =="NaN" || value=="undefined" {
        return nil, nil
    }
    intValue, err := strconv.Atoi(value)
    if err != nil {
        return nil, err
    }
    return &intValue, nil
}

func parseNullableFloat(value string) (*float32, error){
	if value == "" || value =="NaN"  || value =="undefined" {
		return nil, nil
	} 
	floatValue, err := strconv.ParseFloat(value, 32)
	if err !=nil {
		return nil, err
	}
	float32Value := float32(floatValue)
	return &float32Value, nil
}
// Helper function to convert *float64 to *float32
// func float64ToFloat32Ptr(value *float64) *float32 {
// 	if value == nil {
// 		return nil
// 	}
// 	float32Value := float32(*value)
// 	return &float32Value
// }

// Helper function to parse int from string
// func parseInt(value string, fieldName string) (int, error) {
// 	parsedValue, err := strconv.Atoi(value)
// 	if err != nil {
// 		return 0, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid %s", fieldName))
// 	}
// 	return parsedValue, nil
// }

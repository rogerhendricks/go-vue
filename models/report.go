package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	ReportDate string `json:"report_date"`
	FilePath   string `json:"file_path"`
	PatientID uint `json:"patient_id"`
	Patient Patient `gorm:"foreignKey:PatientID"`
}
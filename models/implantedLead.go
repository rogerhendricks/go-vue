package models

import "gorm.io/gorm"

type ImplantedLead struct {
    gorm.Model
    ImplantDate string `json:"implant_date"`
    ExplantDate string `json:"explant_date"`
	IsActive string `json:"is_active"`	
    LeadID uint `json:"lead_id"`
    DoctorID uint `json:"doctor_id"`
    PatientID uint `json:"patient_id"`
    Lead Lead `gorm:"foreignKey:LeadID"`
    Doctor Doctor `gorm:"foreignKey:DoctorID"`
    Patient Patient `gorm:"foreignKey:PatientID"`
}
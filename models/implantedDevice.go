package models

import "gorm.io/gorm"

type ImplantedDevice struct {
    gorm.Model
    ImplantDate string `json:"implant_date"`
    ExplantDate string `json:"explant_date"`
    DeviceID uint `json:"device_id"`
    DoctorID uint `json:"doctor_id"`
    PatientID uint `json:"patient_id"`
    Device Device `gorm:"foreignKey:DeviceID"`
    Doctor Doctor `gorm:"foreignKey:DoctorID"`
    Patient Patient `gorm:"foreignKey:PatientID"`
}
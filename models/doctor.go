package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Addresses []Address `gorm:"foreignKey:DoctorID"`
	PatientID uint
}
type Address struct {
    gorm.Model
    Street   string `json:"street"`
    State    string `json:"state"`
    City     string `json:"city"`
    Zip      string `json:"zip"`
    Country  string `json:"country"`
    DoctorID uint
}

package models

import "gorm.io/gorm"

type Patient struct {
    gorm.Model
    Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	State string `json:"state"`
	City string `json:"city"`
	Zip string `json:"zip"`
	Country string `json:"country"`
	Sex string `json:"sex"`
	Dob string `json:"dob"`
	Active bool `json:"active"`
	
    Doctors []Doctor `gorm:"foreignKey:PatientID"`
}
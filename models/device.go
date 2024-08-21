package models

import "gorm.io/gorm"

type Device struct {
    gorm.Model
	Manufacturer string `json:"manufacturer"`
	DeviceModel string `json:"deviceModel"`
    Name string `json:"name"`
    Type string `json:"type"`
	Header string `json:"header"`
	HasHazard bool `json:"hasHazard"`
	IsMri bool `json:"isMri"`
	Comment string `json:"comment"`
}
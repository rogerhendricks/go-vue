package models

import "gorm.io/gorm"

type Lead struct {
	gorm.Model
	Manufacturer string `json:"manufacturer"`
	LeadModel string `json:"leadModel"`
    Name string `json:"name"`
    Polarity string `json:"polarity"`
	HasHazard bool `json:"hasHazard"`
	IsMri bool `json:"isMri"`
	Comment string `json:"comment"`
}
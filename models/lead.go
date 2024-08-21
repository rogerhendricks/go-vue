package models

import "gorm.io/gorm"

type Lead struct {
	gorm.Model
	Manufacturer string `json:"manufacturer"`
	LeadModel string `json:"deviceModel"`
    Name string `json:"name"`
    Polarity string `json:"type"`
	HasHazard bool `json:"hasHazard"`
	IsMri bool `json:"isMri"`
	Comment string `json:"comment"`
}
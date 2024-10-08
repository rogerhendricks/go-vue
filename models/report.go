package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	// Report Data
	ReportDate   string `json:"report_date"`
	ReportType   string `json:"report_type"`
	ReportStatus string `json:"report_status"`
	// Patient substrate
	CurrentHeartRate  *int    `json:"current_heart_rate"`
	CurrentRhythm     string `json:"current_rhythm"`
	CurrentDependency string `json:"current_dependency"`
	// Patient Arrhythmias
	Mdc_idc__stat_ataf_burden_percent *int `json:"mdc_idc__stat_ataf_burden_percent"`
	// Mdc_idc_stat_pvc_count int `json:"mdc_idc_stat_pvc_count"`
	// Mdc_idc_stat_nsvt_count int `json:"mdc_idc_stat_nsvt_count"`
	// Device Settings
	Mdc_idc_set_brady_mode              string `json:"mdc_idc_set_brady_mode"`
	Mdc_idc_set_brady_lowrate           *int    `json:"mdc_idc_set_brady_lowrate"`
	Mdc_idc_set_brady_max_tracking_rate *int    `json:"mdc_idc_set_brady_max_tracking_rate"`
	Mdc_idc_set_brady_max_sensor_rate   *int    `json:"mdc_idc_set_brady_max_sensor_rate"`
	Mdc_idc_dev_sav                     string `json:"mdc_idc_dev_sav"`
	Mdc_idc_dev_pav                     string `json:"mdc_idc_dev_pav"`
	// Pacing Percentages
	Mdc_idc_stat_brady_ra_percent_paced  *int `json:"mdc_idc_stat_brady_ra_percent_paced"`
	Mdc_idc_stat_brady_rv_percent_paced  *int `json:"mdc_idc_stat_brady_rv_percent_paced"`
	Mdc_idc_stat_brady_lv_percent_paced  *int `json:"mdc_idc_stat_brady_lv_percent_paced"`
	Mdc_idc_stat_brady_biv_percent_paced *int `json:"biv_percent_paced"`
	// Device Diagnostics
	Mdc_idc_batt_volt       *float32    `json:"mdc_idc_batt_volt"`
	Mdc_idc_batt_remaining  string    `json:"mdc_idc_batt_remaining"`
	Mdc_idc_batt_status     string `json:"mdc_idc_batt_status"`
	Mdc_idc_cap_charge_time *float32    `json:"mdc_idc_cap_charge_time"`
	// Ra measurements
	Mdc_idc_msmt_ra_impedance_mean *float32 `json:"mdc_idc_msmt_ra_impedance_mean"`
	Mdc_idc_msmt_ra_sensing_mean   *float32 `json:"mdc_idc_msmt_ra_sensing_mean"`
	Mdc_idc_msmt_ra_threshold      *float32 `json:"mdc_idc_msmt_ra_threshold"`
	Mdc_idc_msmt_ra_pw             *float32 `json:"mdc_idc_msmt_ra_pw"`
	// Rv measurements
	Mdc_idc_msmt_rv_impedance_mean *float32 `json:"mdc_idc_msmt_rv_impedance_mean"`
	Mdc_idc_msmt_rv_sensing_mean   *float32 `json:"mdc_idc_msmt_rv_sensing_mean"`
	Mdc_idc_msmt_rv_threshold      *float32 `json:"mdc_idc_msmt_rv_threshold"`
	Mdc_idc_msmt_rv_pw             *float32 `json:"mdc_idc_msmt_rv_pw"`
	// HV measurements
	Mdc_idc_msmt_shock_impedance *float32 `json:"mdc_idc_msmt_shock_impedance"`
	// Lv measurements
	Mdc_idc_msmt_lv_impedance_mean *float32 `json:"mdc_idc_msmt_lv_impedance_mean"`
	Mdc_idc_msmt_lv_threshold      *float32 `json:"mdc_idc_msmt_lv_threshold"`
	Mdc_idc_msmt_lv_pw             *float32 `json:"mdc_idc_msmt_lv_pw"`
	// Mdc_idc_msmt_lv_sensing_mean int `json:"mdc_idc_msmt_lv_sensing_mean"`

	// Reporter Information
	Comments    string `json:"comments"`
	IsCompleted bool   `json:"is_completed"`
	AuthorID    uint   `json:"author_id"`
	Author      User   `gorm:"foreignKey:AuthorID"`
	// File Path
	FilePath string `json:"file_path"`
	// Patient Information
	PatientID uint    `json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID"`
}

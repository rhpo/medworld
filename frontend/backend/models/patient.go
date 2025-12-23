package models

import (
	"time"
)

// Patient extends User with patient-specific fields
type Patient struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserID           uint      `gorm:"uniqueIndex;not null" json:"userId"`
	User             *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	EmergencyContact string    `json:"emergencyContact"`
	BloodType        string    `json:"bloodType"` // A+, A-, B+, B-, AB+, AB-, O+, O-
	Weight           float64   `json:"weight"`
	MedicalHistory   string    `gorm:"type:text" json:"medicalHistory"` // JSON array
	Allergies        string    `gorm:"type:text" json:"allergies"`      // JSON array
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func (Patient) TableName() string {
	return "patients"
}

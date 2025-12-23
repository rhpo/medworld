package models

import (
	"time"
)

// Doctor extends User with doctor-specific fields
type Doctor struct {
	ID                   uint       `gorm:"primaryKey" json:"id"`
	UserID               uint       `gorm:"uniqueIndex;not null" json:"userId"`
	User                 *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Speciality           string     `json:"speciality"`
	CareerStart          *time.Time `json:"careerStart"`
	ConsultationPrice    float64    `json:"consultationPrice"`
	ConsultationDuration int        `json:"consultationDuration"` // in minutes
	CabinetID            uint       `gorm:"index" json:"cabinetId"`
	Cabinet              *Cabinet   `gorm:"foreignKey:CabinetID" json:"cabinet,omitempty"`
	CreatedAt            time.Time  `json:"createdAt"`
	UpdatedAt            time.Time  `json:"updatedAt"`
}

func (Doctor) TableName() string {
	return "doctors"
}

// GetYearsOfExperience calculates years of experience
func (d *Doctor) GetYearsOfExperience() int {
	if d.CareerStart == nil {
		return 0
	}
	return int(time.Since(*d.CareerStart).Hours() / 24 / 365)
}

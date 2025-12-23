package models

import (
	"time"
)

// Consultation represents a medical consultation record
type Consultation struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	DoctorID      uint         `gorm:"not null;index" json:"doctorId" validate:"required"`
	Doctor        *Doctor      `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	PatientID     uint         `gorm:"not null;index" json:"patientId" validate:"required"`
	Patient       *Patient     `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	AppointmentID uint         `gorm:"index" json:"appointmentId"`
	Appointment   *Appointment `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
	Notes         string       `gorm:"type:text" json:"notes"`
	Prescriptions string       `gorm:"type:text" json:"prescriptions"` // JSON array
	Attachments   string       `gorm:"type:text" json:"attachments"`   // JSON array of file URLs
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     time.Time    `json:"updatedAt"`
}

func (Consultation) TableName() string {
	return "consultations"
}

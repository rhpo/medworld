package models

import (
	"time"
)

type AppointmentStatus string

const (
	AppointmentStatusScheduled  AppointmentStatus = "SCHEDULED"
	AppointmentStatusConfirmed  AppointmentStatus = "CONFIRMED"
	AppointmentStatusInProgress AppointmentStatus = "IN_PROGRESS"
	AppointmentStatusCompleted  AppointmentStatus = "COMPLETED"
	AppointmentStatusCancelled  AppointmentStatus = "CANCELLED"
	AppointmentStatusNoShow     AppointmentStatus = "NO_SHOW"
)

// Appointment represents a medical appointment
type Appointment struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	Date      time.Time         `gorm:"not null" json:"date" validate:"required"`
	Status    AppointmentStatus `gorm:"not null;default:'SCHEDULED'" json:"status"`
	PatientID uint              `gorm:"not null;index" json:"patientId" validate:"required"`
	Patient   *Patient          `gorm:"foreignKey:PatientID" json:"patient,omitempty"`
	DoctorID  uint              `gorm:"not null;index" json:"doctorId" validate:"required"`
	Doctor    *Doctor           `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	CabinetID uint              `gorm:"not null;index" json:"cabinetId" validate:"required"`
	Cabinet   *Cabinet          `gorm:"foreignKey:CabinetID" json:"cabinet,omitempty"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func (Appointment) TableName() string {
	return "appointments"
}

package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Calendar struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	DoctorID     uint           `gorm:"index;not null" json:"doctorId"`
	CabinetID    uint           `json:"cabinetId"`
	Doctor       *Doctor        `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	Cabinet      *Cabinet       `gorm:"foreignKey:CabinetID" json:"cabinet,omitempty"`
	Availability string         `gorm:"type:text" json:"-"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Calendar) TableName() string {
	return "calendars"
}

// Helper to handle JSON marshaling for frontend
func (c *Calendar) MarshalJSON() ([]byte, error) {
	type Alias Calendar
	var slots interface{}
	if c.Availability != "" {
		_ = json.Unmarshal([]byte(c.Availability), &slots)
	}

	return json.Marshal(&struct {
		*Alias
		Availability interface{} `json:"availability"`
	}{
		Alias:        (*Alias)(c),
		Availability: slots,
	})
}

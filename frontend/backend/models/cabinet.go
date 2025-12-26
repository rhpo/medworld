package models

import (
	"time"
)

// Cabinet represents a medical cabinet/clinic
type Cabinet struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"not null" json:"name" validate:"required"`
	Phone            string    `json:"phone"`
	Image            string    `json:"image"`
	AdminID          uint      `gorm:"index" json:"adminId"`
	Admin            *User     `gorm:"foreignKey:AdminID" json:"admin,omitempty"`
	Location         string    `gorm:"type:text" json:"location"`     // JSON: {address, latitude, longitude}
	OpeningHours     string    `gorm:"type:text" json:"openingHours"` // JSON: {day: {open, close}}
	AccessHandicap   *bool     `json:"accessHandicap"`
	HasParking       *bool     `json:"hasParking,omitempty"`
	HasWifi          *bool     `json:"hasWifi,omitempty"`
	AcceptsUrgent    *bool     `json:"acceptsUrgent,omitempty"`
	AcceptsInsurance *bool     `json:"acceptsInsurance,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func (Cabinet) TableName() string {
	return "cabinets"
}

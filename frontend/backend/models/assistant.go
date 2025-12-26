package models

import (
	"time"
)

// Assistant extends User with assistant-specific fields
type Assistant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"uniqueIndex;not null" json:"userId"`
	User      *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	DoctorID  uint      `gorm:"index" json:"doctorId"`
	Doctor    *Doctor   `gorm:"foreignKey:DoctorID" json:"doctor,omitempty"`
	CabinetID uint      `gorm:"index" json:"cabinetId"`
	Cabinet   *Cabinet  `gorm:"foreignKey:CabinetID" json:"cabinet,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Assistant) TableName() string {
	return "assistants"
}

package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserType string

const (
	UserTypeSuperAdmin UserType = "superadmin"
	UserTypeAdmin      UserType = "admin"
	UserTypeDoctor     UserType = "doctor"
	UserTypeAssistant  UserType = "assistant"
	UserTypePatient    UserType = "patient"
)

// User base model for all user types
type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	FirstName   string         `gorm:"not null" json:"firstName" validate:"required"`
	LastName    string         `gorm:"not null" json:"lastName" validate:"required"`
	Email       string         `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Password    string         `gorm:"not null" json:"-"`
	PhoneNumber string         `json:"phoneNumber"`
	AvatarURL   string         `json:"avatarUrl"`
	Address     string         `json:"address"`
	Gender      string         `json:"gender"` // male, female
	DateOfBirth *time.Time     `json:"dateOfBirth"`
	Type        UserType       `gorm:"not null;index" json:"type" validate:"required"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeSave hook to hash password
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword verifies the password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GetFullName returns the full name
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

// TableName specifies the table name
func (User) TableName() string {
	return "users"
}

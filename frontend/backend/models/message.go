package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	SenderID   uint           `gorm:"index;not null" json:"senderId"`
	Sender     *User          `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	ReceiverID uint           `gorm:"index;not null" json:"receiverId"`
	Receiver   *User          `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	IsRead     bool           `gorm:"default:false" json:"isRead"`
	CreatedAt  time.Time      `json:"createdAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Message) TableName() string {
	return "messages"
}

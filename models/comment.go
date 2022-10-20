package models

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id" gorm:"foreignKey:ID;"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message" gorm:"type varchar(191)"`
	CreatedAt time.Time `json:"created_at" gorm:"type datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type datetime"`
}

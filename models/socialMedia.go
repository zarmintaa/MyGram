package models

import "time"

type SocialMedia struct {
	Id             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `json:"name" gorm:"type varchar(50); not null" validate:"required"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"type varchar(191); not null" validate:"required"`
	UserId         uint      `json:"user_id" validate:"required"`
	CreatedAt      time.Time `json:"created_at" gorm:"type datetime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"type datetime"`
}

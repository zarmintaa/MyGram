package models

import "time"

type SocialMedia struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name" gorm:"type varchar(50); not null" validate:"required"`
	Social_media_url string    `json:"social_media_url" gorm:"type varchar(191); not null" validate:"required"`
	User_id          uint      `json:"user_id" validate:"required"`
	Created_at       time.Time `json:"created_at" gorm:"type datetime"`
	Updated_at       time.Time `json:"updated_at" gorm:"type datetime"`
}

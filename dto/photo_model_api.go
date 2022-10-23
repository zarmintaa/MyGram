package dto

import "time"

type PhotoRequest struct {
	Title    string `json:"title" gorm:"type varchar(191);not null" validate:"required"`
	Caption  string `json:"caption" gorm:"type varchar(191);not null" validate:"required"`
	PhotoUrl string `json:"photo_url" gorm:"type varchar(191);not null" validate:"required"`
}

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username" `
	Email    string `json:"email" `
}

type PhotoResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user"`
}

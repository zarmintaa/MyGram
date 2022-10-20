package utils

import "time"

type CommentRequest struct {
	Message string `json:"message" validate:"required,max=191"`
	PhotoId uint   `json:"photo_id" validate:"required"`
}

type UpdateCommentMessage struct {
	Message string `json:"message" validate:"required"`
}

type Photo struct {
	Id       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   uint   `json:"user_id"`
}
type CommentResponse struct {
	Id        string    `json:"id"`
	Message   string    `json:"message"`
	PhotoId   string    `json:"photo_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user"`
	Photo     *Photo    `json:"photo"`
}

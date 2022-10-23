package dto

type SocialMediaRequest struct {
	Name           string `json:"name" validate:"required,max=50"`
	SocialMediaUrl string `json:"social_media_url" validate:"required,max=191"`
}

type SocialMediaResponse struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	UserId         string `json:"-"`
	User           *User  `json:"user"`
}

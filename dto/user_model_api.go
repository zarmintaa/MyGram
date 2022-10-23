package dto

type RegisterRequest struct {
	Username string `json:"username" gorm:"type varchar(10);unique;not null" validate:"required,lte=100"`
	Email    string `json:"email" gorm:"type varchar(191);not null;unique" validate:"required,lte=100"`
	Password string `json:"password"  gorm:"type varchar(191); not null" validate:"required,lte=100"`
	Age      int    `json:"age" validate:"required,lte=100,gte=8"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"type varchar(191);not null;unique" validate:"required,lte=100,gte=8"`
	Password string `json:"password"  gorm:"type varchar(191); not null" validate:"required,lte=100,gte=8"`
}

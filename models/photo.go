package models

import "time"

type Photo struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title" gorm:"type varchar(50); not null"`
	Caption    string    `json:"caption" gorm:"type text"`
	Photo_url  string    `json:"photo_url" gorm:"type varchar(191); not null"`
	User_id    uint      `json:"user_id"`
	Comment    []Comment `gorm:"foreignKey Photo_Id; references Id;constraint:onDelete:CASCADE" json:"comments"`
	Created_at time.Time `json:"created_at" gorm:"type datetime"`
	Updated_at time.Time `json:"updated_at" gorm:"type datetime"`
}

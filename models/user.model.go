package models

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" form:"name" validate:"gte=6,lte=32" gorm:"not null"`
	Email string `json:"email" form:"email" validate:"required,email" gorm:"not null"`
}

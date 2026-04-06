package models

type User struct {
	ID        int64   `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
	AvatarURL *string `json:"avatar_url"`
}

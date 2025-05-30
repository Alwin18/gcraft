package models

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

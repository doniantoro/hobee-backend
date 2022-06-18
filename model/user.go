package model

import "time"

type User struct {
	Id              uint      `json:"id"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	RegisterMethode int       `json:"register_methode"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleteted_at"`
}

package request

import "time"

type User struct {
	CreatedAt       time.Time
	Email           string `json:"email" `
	Password        string `json:"password" binding:"required"`
	RegisterMethode int    `json:"register_methode" binding:"required,number"`
}
type UserInput struct {
	CreatedAt time.Time
	Email     string `json:"email,omitempty" `
	Password  string `json:"password,omitempty"  `
}

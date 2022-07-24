package model

import "time"

type User struct {
	Id              uint      `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Phone           string    `json:"Phone"`
	Gender          int       `json:"gender"`
	Birthday        time.Time `json:"birthday"`
	role        	int 	`json:"role"`
	RegisterMethode int       `json:"register_methode"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleteted_at"`
}

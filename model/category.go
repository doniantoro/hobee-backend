package model

import "time"


type Category struct {
	Id        	uint      `json:"id"`
	Name   		string    `json:"name"`
	Icon       	string    `json:"icon"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	DeletedAt 	time.Time `json:"deleteted_at"`
}

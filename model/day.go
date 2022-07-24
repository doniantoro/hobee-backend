package model

import "time"


type Day struct {
	Id        	uint      `json:"id"`
	Name   		string    `json:"name"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	DeletedAt 	time.Time `json:"deleteted_at"`
}

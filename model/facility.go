package model

import "time"


type Facility struct {
	Id        uint      `json:"id"`
	VenueId   int    `json:"venue_id"`
	Light       bool    `json:"light"`
	Shower       bool    `json:"shower"`
	Store       bool    `json:"store"`
	Toilet       bool    `json:"toilet"`
	Wash       bool    `json:"day"`
	Food       bool    `json:"food"`
	Drink       bool    `json:"drink"`
	Chair       bool    `json:"chair"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleteted_at"`
}

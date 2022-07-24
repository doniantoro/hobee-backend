package model

import "time"

type Review struct {
	Id          uint      `json:"id"`
	VenueId     string    `json:"venue_id"`
	UserId      string    `json:"user_id"`
	rating      int       `json:"rating"`
	description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleteted_at"`
}

package model

import "time"


type Price struct {
	Id        uint      `json:"id"`
	VenueId   string    `json:"venue_id"`
	Day       string    `json:"day"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleteted_at"`
}

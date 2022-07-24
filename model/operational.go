package model

import "time"


type Operational struct {
	Id        uint      `json:"id"`
	VenueId   int    	`json:"venue_id"`
	DayId     string    `json:"day_id"`
	StartTime string    `json:"start_time"`
	EndTime string    	`json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleteted_at"`
}

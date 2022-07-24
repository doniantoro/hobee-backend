package model

import "time"


type ReviewAttachment struct {
	Id        uint      `json:"id"`
	ReviewID   int    	`json:"review_id"`
	url     string    	`json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleteted_at"`
}

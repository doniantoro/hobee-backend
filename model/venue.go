package model

import "time"

type Venue struct {
	Id            uint      `json:"id"`
	UserId        string    `json:"user_id"`
	CategoryId    string    `json:"category_id"`
	FacilitityId  int       `json:"facility_id"`
	OperationalId int       `json:"operational_id"`
	DistrictId    int       `json:"district_id"`
	Description   string       `json:"description"`
	FullAddress   string    `json:"full_address"`
	Lat           string    `json:"lat"`
	Long          string    `json:"long"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleteted_at"`
}

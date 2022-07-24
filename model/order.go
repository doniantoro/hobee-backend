package model

import "time"

type Order struct {
	Id          uint      	`json:"id"`
	VenueId     int    	`json:"venue_id"`
	BookingId    int   `json:"user_id"`
	VoucherID    int      `json:"voucher_id"`
	PaymentType  string    	`json:"payment_type"`
	TotalPayment int    	`json:"total_payment"`
	Status 		 string    		`json:"status"`
	BookingCode  string    	`json:"booking_code"`
	Date 		string    			`json:"date"`
	StartTime 		string    			`json:"start_time"`
	EndTime 		string    			`json:"end_time"`
	CreatedAt   time.Time 	`json:"created_at"`
	UpdatedAt   time.Time 	`json:"updated_at"`
	DeletedAt   time.Time 	`json:"deleteted_at"`
}

package model

type Booking struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Destination string  `json:"destination"`
	Price       float64 `json:"price"`
}

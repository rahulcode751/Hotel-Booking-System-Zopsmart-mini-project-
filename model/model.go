package model

type Booking struct {
	ID           int    `json:"id"`
	HotelName    string `json:"hotelname"`
	CustomerName string `json:"customername"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
	CustomerContact      string `json:"customercontact"`
}

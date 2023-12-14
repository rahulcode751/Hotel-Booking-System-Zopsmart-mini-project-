package model

type Student struct {
	ID           int    `json:"id"`
	HotelName    string `json:"hotelname"`
	CustomerName string `json:"customername"`
	Date         string `json:"date"`
	Price        int    `json:"price"`
}

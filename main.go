package main

import (
	"zopsmart-mini-project/datastore"
	"zopsmart-mini-project/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	/*
		Project By Rahul Bairagi (201B204) JUET
		+91-6264959991
		Hotel Booking System
	*/
	app := gofr.New()
	DStore := datastore.New()
	handleData := handler.New(DStore)
	// Setting Server PORT
	app.Server.HTTP.Port = 8000
	app.GET("/hotelbooking/detail/{id}", handleData.GetByID)
	app.GET("/hotelbooking/details", handleData.GetAllBookings)
	app.POST("/hotelbooking/add", handleData.Create)
	app.PUT("/hotelbooking/update/{id}", handleData.Update)
	app.DELETE("/hotelbooking/delete/{id}", handleData.Delete)
	// Starting Server
	app.Start()
}

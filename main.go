package main

import (
	"zopsmart-mini-project/datastore"
	"zopsmart-mini-project/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	// Project By Rahul bairagi(201B204) +91-6264959991 Hotel Booking System
	app := gofr.New()
	DStore := datastore.New()
	handleData := handler.New(DStore)
	// Setting Server PORT
	app.Server.HTTP.Port = 8000
	app.GET("/hotelbooking/{id}", handleData.GetByID)
	app.POST("/hotelbooking", handleData.Create)
	app.PUT("/hotelbooking/{id}", handleData.Update)
	app.DELETE("/hotelbooking/{id}", handleData.Delete)
	// Starting Server
	app.Start()
}

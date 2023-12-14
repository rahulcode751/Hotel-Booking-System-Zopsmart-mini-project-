package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()
	app.Server.HTTP.Port = 8000
	app.Start()
}

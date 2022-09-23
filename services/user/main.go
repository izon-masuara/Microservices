package main

import (
	"net/http"
	"user/config"
	"user/routers"
)

func main() {
	config.Connect()
	routers.Home()
	address := "localhost:3000"
	err := http.ListenAndServe(address, nil)
	// Penerapan middleware error listen
	if err != nil {
		panic("Error connect to server")
	}
}

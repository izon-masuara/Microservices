package main

import (
	"net/http"
	"user/configs"
	"user/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error .env file")
	}
	configs.Connect()
	routers.ApiUser()
	address := "localhost:3000"
	err = http.ListenAndServe(address, nil)
	// Penerapan middleware error listen
	if err != nil {
		panic("Error connect to server")
	}
}

package main

import (
	"net/http"
	"user/package/configs"
	"user/package/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error .env file")
	}
	configs.Connect()
	routers.ApiUser()
	address := "localhost:3000"
	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic("Error connect to server")
	}
}

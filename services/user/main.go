package main

import (
	"net/http"
	"user/package/configs"
	"user/package/routers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic("Error .env file")
	}

	r := mux.NewRouter()
	configs.Connect()
	routers.ApiUser(r)

	http.ListenAndServe(":3000", r)
}

package main

import (
	"net/http"
	"user/routers"
)

func main() {
	routers.Home()

	address := "localhost:3000"

	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic("Error connect to server")
	}
}

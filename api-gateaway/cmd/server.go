package main

import (
	"api-gateaway/cmd/controller"
	"api-gateaway/cmd/db"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	db.Connect()

	router.HandleFunc("/", controller.GetDataFormQuery)
	http.ListenAndServe(":3002", router)
}

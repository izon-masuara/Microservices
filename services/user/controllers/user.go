package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user/models"
	"user/query"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var message string
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			message = "Invalid payload from client"
		}
		defer r.Body.Close()
		var payload models.UserLogin
		json.Unmarshal(res, &payload)
		message, statusCode := query.CheckUserExists(payload)
		if statusCode == 200 {
			w.WriteHeader(http.StatusOK)
		} else if statusCode == 400 {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write([]byte(message))
	}
}

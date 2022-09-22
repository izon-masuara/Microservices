package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user/models"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []UserLogin{
	{
		Username: "Budi",
		Password: "saha234",
	},
	{
		Username: "Jindan",
		Password: "7834j",
	},
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		models.GetAllUser()
		res, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		var payload UserLogin
		json.Unmarshal(res, &payload)
		for _, user := range users {
			if user.Username == payload.Username && user.Password == payload.Password {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Success Login"))
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid Username or Password"))
	}
}

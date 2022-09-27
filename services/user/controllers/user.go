package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user/helpers"
	"user/models"
	"user/query"

	"github.com/golang-jwt/jwt/v4"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var message string
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			message = "Invalid payload from client"
			http.Error(w, message, 400)
		}
		defer r.Body.Close()

		// Get model user
		var payload models.UserLogin
		json.Unmarshal(res, &payload)
		message, statusCode := query.CheckUserExists(payload)
		if statusCode == 202 {
			w.WriteHeader(http.StatusOK)
			var accessToken = models.Token{
				AccessToken: message,
			}
			data, err := json.Marshal(accessToken)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, "Bad request", 400)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		} else if statusCode == 400 {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, message, 400)
		} else if statusCode == 404 {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, message, 400)
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var message string
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			message = "Invalid payload from client"
			http.Error(w, message, 400)
		}
		defer r.Body.Close()

		// Get model user
		var payload models.UserLogin
		json.Unmarshal(res, &payload)
		if len(payload.Username) < 3 || len(payload.Password) < 3 {
			w.WriteHeader(http.StatusBadRequest)
			message = "Username or Password must be more than 3 characters"
		} else {
			msg, statusCode := query.AddUser(payload)
			message = msg
			if statusCode == 201 {
				w.WriteHeader(http.StatusOK)
			} else if statusCode == 400 {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
		var msg = models.GeneralMessage{
			Message: message,
		}
		result, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "Bad request", 400)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(result)
	}
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	type accessToken struct {
		Token string `json:"accessToken"`
	}
	type payload struct {
		User_id  float64 `json:"userId"`
		Username string  `json:"username"`
	}
	var flag = 0
	if r.Method == "POST" {
		res, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			flag = 1
		}
		defer r.Body.Close()
		var token accessToken
		json.Unmarshal(res, &token)
		claims, err := helpers.VerifyToken(token.Token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			flag = 1
		}
		if flag == 0 {
			userId := claims.(jwt.MapClaims)["user_id"].(float64)
			username := claims.(jwt.MapClaims)["username"].(string)
			var payload = payload{
				User_id:  userId,
				Username: username,
			}
			data, err := json.Marshal(payload)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			w.Write(data)
		} else {
			var msg = models.GeneralMessage{
				Message: "Bad request",
			}
			res, err := json.Marshal(msg)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				http.Error(w, "Bad request", 400)
			}
			w.Write(res)
		}
	}
}

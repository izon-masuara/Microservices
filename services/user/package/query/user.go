package query

import (
	"fmt"
	"time"
	"user/package/helpers"
	"user/package/models"
)

/*
Cek user exists on database
if exists return access token
*/
func CheckUserExists(payload models.UserLogin) (string, int) {
	var user models.Users
	if err := models.Db.Where("username= ?", payload.Username).First(&user).Error; err != nil {
		return "Invalid username or password", 404
	}
	_, err := helpers.ComparePass(user.Password, payload.Password)
	if user.Password == "" {
		return "Invalid username or password", 404
	} else if err != nil {
		return "Invalid usernname or password", 404
	}
	accessToken, err := helpers.GetToken(int(user.ID), user.Username)
	if err != nil {
		return "Invalid usernname or password", 404
	}
	return accessToken, 202
}

// Add user if not exists
func AddUser(payload models.UserLogin) (string, int) {
	db := models.Db
	password, err := helpers.HassPass(payload.Password)
	if err != nil {
		return "Failed Register", 400
	}

	var user = models.Users{
		Username:   payload.Username,
		Password:   fmt.Sprintf("%s", password),
		Last_login: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err.Error())
		msg := err.Error()
		if msg == `pq: duplicate key value violates unique constraint "users_username_key"` {
			return "username already exists", 400
		}
		return "Failed Register", 400
	}
	return "Success Register", 201
}

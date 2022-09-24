package query

import (
	"time"
	"user/configs"
	"user/helpers"
	"user/models"
)

func CheckUserExists(payload models.UserLogin) (string, int) {
	db := configs.Db
	res, err := db.Query(`
		SELECT "user_id", "username", "password"
		FROM "users"
		WHERE username=$1
	`, payload.Username)
	if err != nil {
		return "Data not found", 404
	}
	defer res.Close()
	var user models.User
	for res.Next() {
		err = res.Scan(&user.User_id, &user.Username, &user.Password)
		if err != nil {
			return "Invalid usernname or password", 400
		}
	}
	_, err = helpers.ComparePass(user.Password, payload.Password)
	if user.Password == "" {
		return "Invalid usernname or password", 400
	} else if err != nil {
		return "Invalid usernname or password", 400
	}
	accessToken, err := helpers.GetToken(user.User_id, user.Username)
	if err != nil {
		panic(err.Error())
	}
	return accessToken, 200
}

func AddUser(payload models.UserLogin) (string, int) {
	db := configs.Db
	password := helpers.HassPass(payload.Password)
	_, err := db.Query(`
		INSERT INTO "users" ("username","password","created_at","last_login")
		VALUES ($1,$2,$3,$4)
	`, payload.Username, password, time.Now(), time.Now())
	if err != nil {
		msg := err.Error()
		if msg == `pq: duplicate key value violates unique constraint "users_username_key"` {
			return "username already exists", 400
		}
		return "Failed Register", 400
	}
	return "Success Register", 201
}

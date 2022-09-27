package query

import (
	"time"
	"user/package/helpers"
	"user/package/models"
)

/*
Cek user exists on database
if exists return access token
*/
func CheckUserExists(payload models.UserLogin) (string, int) {
	db := models.Db
	res, err := db.Query(`
		SELECT "user_id", "username", "password"
		FROM "users"
		WHERE username=$1
	`, payload.Username)
	if err != nil {
		return "Invalid username or password", 404
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
		return "Invalid username or password", 404
	} else if err != nil {
		return "Invalid usernname or password", 404
	}
	accessToken, err := helpers.GetToken(user.User_id, user.Username)
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
	_, err = db.Query(`
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

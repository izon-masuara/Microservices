package query

import (
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
		panic(err.Error())
	}
	defer res.Close()
	var user models.User
	for res.Next() {
		err = res.Scan(&user.User_id, &user.Username, &user.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	if user.Password == "" {
		return "Invalid usernname or password", 400
	} else if payload.Password != user.Password {
		return "Invalid usernname or password", 400
	}
	accessToken, err := helpers.GetToken(user.User_id, user.Username)
	if err != nil {
		panic(err.Error())
	}
	return accessToken, 200
}

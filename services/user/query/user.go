package query

import (
	"user/config"
	"user/models"
)

func SuccessLogin(payload models.UserLogin) (string, int) {
	db := config.Db
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
	return "Access Token", 200
}

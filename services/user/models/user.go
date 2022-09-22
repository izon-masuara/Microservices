package models

import (
	"fmt"
	"user/config"
)

func GetAllUser() {
	err := config.Db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Masuk di models")
}

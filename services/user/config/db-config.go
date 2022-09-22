package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST    = "localhost"
	PORT    = 5432
	USER    = "postgres"
	PASS    = "wppq"
	DB_NAME = "User_Movie"
)

var Db *sql.DB

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASS, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected!")
	Db = db
}

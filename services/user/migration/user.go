package main

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

func Connect() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASS, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected!")
	return db
}

func main() {
	db := Connect()
	queryCreateTableUser := `
	CREATE TABLE users (
		user_id serial primary key,
		username varchar(40) unique not null,
		password varchar(30) not null,
		created_at timestamp not null,
		last_login timestamp
	);
	`
	_, err := db.Exec(queryCreateTableUser)
	if err != nil {
		panic(err.Error())
	}
}

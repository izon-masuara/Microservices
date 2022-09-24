package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect(args []string) *sql.DB {
	HOST := "localhost"
	PORT := 5432
	USER := args[0]
	PASS := args[1]
	DB_NAME := "User_Movie"
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASS, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected!")
	return db
}

func main() {
	args := os.Args[1:]
	db := Connect(args)
	queryCreateTableUser := `
	CREATE TABLE users (
		user_id serial primary key,
		username varchar(40) unique not null,
		password varchar(100) not null,
		created_at timestamp not null,
		last_login timestamp
	);
	`
	_, err := db.Exec(queryCreateTableUser)
	if err != nil {
		panic(err.Error())
	}
}

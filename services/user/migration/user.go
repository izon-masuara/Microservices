package main

import (
	"database/sql"
	"fmt"
	"os"
	"user/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect(args []string) *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error .env file")
	}
	appConfig := models.EnvConfig{
		HOST:    os.Getenv("DB_HOST"),
		PORT:    os.Getenv("DB_PORT"),
		USER:    os.Getenv("DB_USER"),
		PASS:    os.Getenv("DB_PASS"),
		DB_NAME: os.Getenv("DB_NAME"),
	}
	psqlconn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		appConfig.HOST, appConfig.PORT, appConfig.USER, appConfig.PASS, appConfig.DB_NAME)
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

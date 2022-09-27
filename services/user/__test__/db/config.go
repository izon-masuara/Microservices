package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"user/package/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type EnvConfigTest struct {
	HOST, PORT, USER, PASS, DB_NAME_TEST string
}

func Connect() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error .env file")
	}
	appConfig := EnvConfigTest{
		HOST:         os.Getenv("DB_HOST"),
		PORT:         os.Getenv("DB_PORT"),
		USER:         os.Getenv("DB_USER"),
		PASS:         os.Getenv("DB_PASS"),
		DB_NAME_TEST: os.Getenv("DB_NAME_TEST"),
	}
	psqlconn := fmt.Sprintf(
		`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		appConfig.HOST, appConfig.PORT, appConfig.USER, appConfig.PASS, appConfig.DB_NAME_TEST)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected!")
	models.Db = db
}

func Migrate() {
	Connect()
	db := models.Db
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

func DropTable() {
	Connect()
	db := models.Db
	queryCreateTableUser := `DROP TABLE "users"`
	_, err := db.Exec(queryCreateTableUser)
	if err != nil {
		panic(err.Error())
	}
}

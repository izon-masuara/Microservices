package configs

import (
	"database/sql"
	"fmt"
	"os"
	"user/models"

	_ "github.com/lib/pq"
)

var Db *sql.DB

// return error yang nantinya akan di tangkap middleware jika error
func Connect() {
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
	Db = db
}

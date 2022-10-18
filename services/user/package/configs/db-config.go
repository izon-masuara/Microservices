package configs

import (
	"fmt"
	"os"
	"user/package/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to database
func Connect() {
	appConfig := models.EnvConfig{
		HOST:    os.Getenv("DB_HOST"),
		PORT:    os.Getenv("DB_PORT"),
		USER:    os.Getenv("DB_USER"),
		PASS:    os.Getenv("DB_PASS"),
		DB_NAME: os.Getenv("DB_NAME"),
	}
	psqlconn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		appConfig.USER, appConfig.PASS, appConfig.HOST, appConfig.PORT, appConfig.DB_NAME)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if err := db.AutoMigrate(models.Users{}); err != nil {
		panic(err)
	}

	models.Db = db
}

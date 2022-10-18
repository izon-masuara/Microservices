package models

import (
	"gorm.io/gorm"
)

type EnvConfig struct {
	HOST, PORT, USER, PASS, DB_NAME string
}

var Db *gorm.DB

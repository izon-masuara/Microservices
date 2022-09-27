package models

import "database/sql"

type EnvConfig struct {
	HOST, PORT, USER, PASS, DB_NAME string
}

var Db *sql.DB

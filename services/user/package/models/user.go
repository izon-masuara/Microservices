package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"serial;primary_key" binding:"required"`
	Username   string    `json:"username" gorm:"unique;not null" binding:"required"`
	Password   string    `json:"password" gorm:"not null" binding:"required"`
	Last_login time.Time `json:"last_login" gorm:"timestamptz" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Username   string `json:"username"`
	FullName   string `json:"full_name"`
	KTPID      string `json:"ktp_id"`
	Role       string `json:"role"`
	Address    string `json:"address"`
	Password   string `json:"-"`
}

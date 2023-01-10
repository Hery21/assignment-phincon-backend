package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model `json:"-"`
	ID         int    `json:"id"`
	Place      string `json:"place"`
	Address    string `json:"address"`
}

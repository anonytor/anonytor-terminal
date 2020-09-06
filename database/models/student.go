package models

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	CardID string `json:"card_id"`
	Password string `json:"password"`
}

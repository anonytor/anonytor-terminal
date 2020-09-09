package models

import (
	"anonytor-terminal/runtime/random"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	TOKEN_LENGTH = 32
)

type Token struct {
	ID        uint      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `gorm:"unique" json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

func CheckToken(db *gorm.DB, token string) bool {
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()
	var t Token
	if v := tx.Where("token = ?", token).First(&t); gorm.IsRecordNotFoundError(v.Error) {
		return false
	} else if v.Error != nil {
		panic(v.Error)
	}
	if time.Now().After(t.ExpiredAt) {
		if v := tx.Delete(&t); v.Error != nil {
			panic(v.Error)
		}
		return false
	}
	if v := tx.Commit(); v.Error != nil {
		panic(v.Error)
	}
	return true
}

func NewToken(db *gorm.DB, expiredAt time.Time) *Token {
	token := Token{
		Token:     random.String(TOKEN_LENGTH, random.AlphaNumeric),
		ExpiredAt: expiredAt,
	}
	if v := db.Create(&token); v.Error != nil {
		panic(v.Error)
	}
	return &token
}

func DeleteToken(db *gorm.DB, token string) {
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()
	var t Token
	if v := tx.Where("token = ?", token).First(&t); v.Error != nil {
		panic(v.Error)
	}
	if v := tx.Delete(&t); v.Error != nil {
		panic(v.Error)
	}
	if v := tx.Commit(); v.Error != nil {
		panic(v.Error)
	}
}

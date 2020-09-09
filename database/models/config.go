package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Config struct {
	ID        uint
	CreatedAt time.Time
	Key       string
	Value     string
}

func (c Config) Get(db *gorm.DB, key string) Config {
	if v := db.Model(Config{}).Where("key = ?", key).First(&c); v.Error != nil {
		panic(v.Error)
	}
	return c
}

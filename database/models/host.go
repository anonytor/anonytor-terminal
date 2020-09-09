package models

import (
	"anonytor-terminal/runtime/random"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Host struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Key       string    `json:"-"`
	Name      string    `json:"name"`
	Addr      string    `json:"addr"`
	OSType    int       `json:"os_type"`
	OSInfo    string    `json:"os_info"`
	LastSeen  time.Time `json:"last_seen"`
}

func NewHost(db *gorm.DB, name string) *Host {
	host := Host{
		ID:   uuid.New().String(),
		Key:  random.String(32, random.AlphaNumeric+random.Symbol),
		Name: name,
	}
	if v := db.Create(&host); v.Error != nil {
		panic(v.Error)
	}
	return &host
}

func GetHostById(db *gorm.DB, id string) *Host {
	var host Host
	if v := db.Where("id = ?", id).First(&host); gorm.IsRecordNotFoundError(v.Error) {
		return nil
	} else if v.Error != nil {
		panic(v.Error)
	}
	return &host
}

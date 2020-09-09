package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"monitor-server-backend/runtime/random"
	"time"
)

type Host struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	Key       string
	Name      string
	Addr      string
	IsNAT     bool
	OSType    int
	OSInfo    string
	LastSeen  time.Time
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

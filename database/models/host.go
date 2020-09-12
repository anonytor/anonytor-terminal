package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	"anonytor-terminal/runtime/random"
)

type Host struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Key       string    `json:"-"`
	Name      string    `json:"name"`
	Addr      string    `json:"addr"`
	OS        int       `json:"os"`
	LastSeen  time.Time `json:"last_seen"`
	Status    int       `json:"status"`
}

func NewHost(db *gorm.DB, name string, OS int) *Host {
	host := Host{
		ID:   uuid.New().String(),
		Key:  random.String(32, random.AlphaNumeric+random.Symbol),
		Name: name,
		OS:   OS,
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

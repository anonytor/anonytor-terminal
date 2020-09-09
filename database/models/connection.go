package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Connection struct {
	ID        uint
	CreatedAt time.Time
	Addr      string
	Host      *Host `gorm:"foreignkey:HostID"`
	HostID    string
}

func NewConnection(db *gorm.DB, addr, hostId, key string) *Connection {
	host := GetHostById(db, hostId)
	if host == nil || host.Key != key {
		return nil
	}
	connection := Connection{
		Addr:   addr,
		Host:   host,
		HostID: host.ID,
	}
	if v := db.Create(&host); v.Error != nil {
		panic(v.Error)
	}
	return &connection
}

func GetConnectionsByHostId(db *gorm.DB, hostId string) (connections []Connection) {
	if v := db.Where("host_id = ?", hostId).Find(&connections); v.Error != nil {
		panic(v.Error)
	}
	return
}

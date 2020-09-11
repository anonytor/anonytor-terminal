package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Connection struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Addr      string    `json:"addr"`
	Host      *Host     `gorm:"foreignkey:HostID" json:"-"`
	HostID    string    `json:"host_id"`
	Type      int       `json:"type"`
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

func DeleteConnection(db *gorm.DB, addr string) {
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()
	var t Token
	if v := tx.Where("addr = ?", addr).First(&t); v.Error != nil {
		panic(v.Error)
	}
	if v := tx.Delete(&t); v.Error != nil {
		panic(v.Error)
	}
	if v := tx.Commit(); v.Error != nil {
		panic(v.Error)
	}
}

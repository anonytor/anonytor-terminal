package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Connection struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Addr      string    `json:"addr"`
	HostID    string    `json:"host_id"`
	Type      int       `json:"type"`
}

func NewConnection(db *gorm.DB, addr, hostId, key string, typ int) *Connection {
	host := GetHostById(db, hostId)
	if host == nil || host.Key != key {
		return nil
	}
	connection := Connection{
		Addr:   addr,
		HostID: host.ID,
		Type:   typ,
	}
	if v := db.Create(&connection); v.Error != nil {
		panic(v.Error)
	}
	host.Addr = addr
	host.LastSeen = time.Now()
	host.Status = 1
	if v := db.Save(host); v.Error != nil {
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

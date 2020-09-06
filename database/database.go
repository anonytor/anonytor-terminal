package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"monitor-server-backend/config"
	"monitor-server-backend/database/models"
)

// InitDatabase inits a database connection to mysql
// all parameters needed will loaded from ./conf
// make sure you init ./conf before calling this function
func InitDatabase(conf config.Database) (db *gorm.DB) {
	var err error
	db, err = gorm.Open(conf.Dialect, conf.Parameter)
	if err != nil {
		errStr := fmt.Sprintf("无法建立与数据库的连接！原因：%v", err)
		log.Fatal(errStr)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)

	MigrateAll(db)
	return
}

func MigrateAll(db *gorm.DB) {
	db.AutoMigrate(&models.Student{})
}

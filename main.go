package main

import (
	"monitor-server-backend/config"
	"monitor-server-backend/control"
	"monitor-server-backend/database"
	"monitor-server-backend/runtime/logger"
)

const ConfigurationPath = "./config.json"

func main() {
	// Init Configurations
	conf := config.InitConfig(ConfigurationPath)
	// Init Logger
	logger.InitLogger(conf.GetLogLevel())
	// Init Database
	db := database.InitDatabase(conf.Database)
	// Init ServerSocket
	socket := control.InitSocket(db, conf.Control.Addr)
	// start listening
	go socket.ListenAndServe()
}

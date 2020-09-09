package main

import (
	"anonytor-terminal/api"
	"anonytor-terminal/config"
	"anonytor-terminal/database"
	"anonytor-terminal/runtime/logger"
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
	//socket := control.InitSocket(db, conf.Control.Addr)
	// start listening
	//go socket.ListenAndServe()
	// Init server
	apiServer := api.NewServer(conf.Api, db)
	apiServer.Start()
	select {}
}

package main

import (
	"anonytor-terminal/config"
	"anonytor-terminal/controller"
	"anonytor-terminal/database"
	"anonytor-terminal/runtime/logger"
	"sync"
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
	ctrl := controller.InitController(db, conf.Control.Addr)
	// start listening
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ctrl.ListenAndServe()
	wg.Wait()
}

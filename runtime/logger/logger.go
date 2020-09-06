package logger

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger(level log.Level) {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(level)
	//logger.WithField()
}
func init(){
	log.SetLevel(log.DebugLevel)
}

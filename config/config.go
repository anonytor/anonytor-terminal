package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const (
	Release = "release"
	Debug   = "debug"
)

type Conf struct {
	RunMode  string   `json:"run_mode"`
	Database Database `json:"database"`
	Api      Api      `json:"api"`
	Control  Control  `json:"control"`
}

func (conf *Conf) GetLogLevel() log.Level {
	if conf.RunMode == Release {
		return log.ErrorLevel
	} else {
		return log.DebugLevel
	}
}

type Database struct {
	Dialect   string `json:"dialect"`
	Parameter string `json:"parameter"`
}

type Control struct {
	Addr string `json:"addr"`
}

type Api struct {
	Addr string `json:"addr"`
}


func InitConfig(path string) *Conf {
	var conf Conf
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Debug("found configuration file at " + path)
	}
	tmp, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Debug("file read successfully")
	}
	err = json.Unmarshal(tmp, &conf)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Debug("parse configuration successfully ")
	}
	return &conf
}

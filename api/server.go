package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "monitor-server-backend/api/middlewares"
	"monitor-server-backend/config"
)

type Server struct {
	gin  *gin.Engine
	db   *gorm.DB
	conf config.Panel
}

func NewServer(conf config.Panel, db *gorm.DB) *Server {
	return &Server{
		db:   db,
		conf: conf,
	}
}

func (s *Server) init() {
	s.gin = gin.New()
	store, _ := redis.NewStore(10, "tcp", s.conf.Redis.Host, s.conf.Redis.Password)
	s.gin.Use(gin.Logger())
	s.gin.Use(Recovery())
	s.gin.Use(sessions.Sessions("session", store))
	s.gin.Use(SetDb(s.db))
	s.gin.Use(CheckSession())
	s.gin.NoMethod(Handler404())
	s.gin.NoRoute(Handler404())
}

func (s *Server) Start() {
	s.init()
	go func() {
		err := s.gin.Run(s.conf.Addr)
		if err != nil {
			panic(err)
		}
	}()
}

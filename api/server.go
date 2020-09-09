package api

import (
	"anonytor-terminal/api/handlers"
	. "anonytor-terminal/api/middlewares"
	"anonytor-terminal/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	gin  *gin.Engine
	db   *gorm.DB
	conf config.Api
}

func NewServer(conf config.Api, db *gorm.DB) *Server {
	return &Server{
		db:   db,
		conf: conf,
	}
}

func (s *Server) init() {
	s.gin = gin.New()
	s.gin.Use(gin.Logger())
	s.gin.Use(Recovery())
	s.gin.Use(SetDb(s.db))
	s.gin.Use(Cors())
	s.gin.Use(Auth())
	s.gin.NoMethod(Handler404())
	s.gin.NoRoute(Handler404())
	handlers.RegisterAgent(s.gin.Group("agent"))
	handlers.RegisterConnection(s.gin.Group("connection"))
	handlers.RegisterHost(s.gin.Group("host"))
	handlers.RegisterInfo(s.gin.Group("info"))
	handlers.RegisterPing(s.gin.Group("ping"))
	handlers.RegisterToken(s.gin.Group("token"))
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

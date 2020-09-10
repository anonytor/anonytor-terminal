package api

import (
	"anonytor-terminal/api/handlers"
	. "anonytor-terminal/api/middlewares"
	"anonytor-terminal/config"
	"anonytor-terminal/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	gin  *gin.Engine
	db   *gorm.DB
	conf config.Api
	ctrl *controller.Controller
}

func NewServer(conf config.Api, db *gorm.DB, ctrl *controller.Controller) *Server {
	return &Server{
		db:   db,
		conf: conf,
		ctrl: ctrl,
	}
}

func (s *Server) init() {
	s.gin = gin.New()
	s.gin.Use(gin.Logger())
	s.gin.Use(Recovery())
	s.gin.Use(SetDb(s.db))
	s.gin.Use(SetController(s.ctrl))
	s.gin.Use(Cors())
	s.gin.Use(Auth())
	s.gin.NoMethod(Handler404())
	s.gin.NoRoute(Handler404())
	handlers.RegisterAgent(s.gin.Group("agent"))
	handlers.RegisterConnection(s.gin.Group("connection"))
	handlers.RegisterHost(s.gin.Group("host"))
	handlers.RegisterInfo(s.gin.Group("info"))
	handlers.RegisterPing(s.gin.Group("ping"))
	handlers.RegisterTask(s.gin.Group("task"))
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

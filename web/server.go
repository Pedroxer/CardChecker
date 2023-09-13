package web

import (
	"github.com/Pedroxer/CardChecker/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config util.Config
}

func NewServer(conf util.Config) *Server {
	server := &Server{
		config: conf,
	}
	server.setupRoutes()
	return server
}

func (serv *Server) setupRoutes() {
	router := gin.Default()
	router.GET("card", serv.validCheck)
	serv.router = router
}

func (serv *Server) Start() error {
	return serv.router.Run(serv.config.ServerAddress)
}

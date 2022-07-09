package api

import (
	"Werken/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := Server{}
	router := gin.Default()

	router.GET("/realtime", handler.GetRealTime())
	router.GET("/auth", handler.GetAuth())
	router.GET("/user", handler.GetUser())
	router.GET("/jobs", handler.GetJobs())
	router.GET("/finder", handler.GetFinder())
	router.GET("/wallet", handler.GetWallet())

	router.POST("/realtime", handler.PostRealTime())
	router.POST("/auth", handler.PostAuth())
	router.POST("/user", handler.PostUser())
	router.POST("/jobs", handler.PostJobs())
	router.POST("/wallet", handler.PostWallet())
	server.router = router
	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

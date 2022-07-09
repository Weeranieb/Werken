package api

import (
	"Werken/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func registerRouter(r *gin.RouterGroup) {
	r.GET("/realtime", handler.GetRealTime())
	r.GET("/auth", handler.GetAuth())
	r.GET("/user", handler.GetUser())
	r.GET("/jobs", handler.GetJobs())
	r.GET("/finder", handler.GetFinder())
	r.GET("/wallet", handler.GetWallet())

	r.POST("/realtime", handler.PostRealTime())
	r.POST("/auth", handler.PostAuth())
	r.POST("/user", handler.PostUser())
	r.POST("/jobs", handler.PostJobs())
	r.POST("/wallet", handler.PostWallet())
}

func NewServer() *gin.Engine {
	router := gin.Default()

	r := router.Group("/api")

	registerRouter(r)

	return router
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

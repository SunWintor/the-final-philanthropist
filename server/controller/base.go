package controller

import (
	"github.com/SunWintor/tfp/configs"
	"net/http"

	"github.com/SunWintor/tfp/server/service"

	"github.com/gin-gonic/gin"
)

var svr *service.Service

func Handle(r *gin.Engine, s *service.Service) *gin.Engine {
	configs.LoadConf()
	svr = s
	g := r.Group("/game/tfp/")
	r.LoadHTMLGlob("web/templates/*")
	handleHtml(g)
	handleHome(g)
	handleAuth(g)
	handleRoom(g)
	handleGame(g)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

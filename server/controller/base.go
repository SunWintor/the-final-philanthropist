package controller

import (
	"net/http"

	"github.com/SunWintor/tfp/server/service"

	"github.com/gin-gonic/gin"
)

var svr *service.Service

func Handle(r *gin.Engine) *gin.Engine {
	svr = service.New()

	handleHome(r)
	handleAuth(r)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

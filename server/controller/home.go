package controller

import (
	"github.com/SunWintor/tfp/server/controller/filter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHome(r *gin.Engine) {
	home := r.Group("/home", filter.JWTAuthHandler())

	home.GET("/info", homeInfo)
}

func homeInfo(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

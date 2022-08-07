package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleHome(r *gin.Engine) {
	home := r.Group("/home")

	home.GET("/info", homeInfo)
}

func homeInfo(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleHtml(r *gin.Engine) {

	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.GET("/room", func(c *gin.Context) {
		c.HTML(http.StatusOK, "room.html", nil)
	})
}

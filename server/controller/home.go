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
	// 获取当前账号信息，比如用户名啥的，展示在页面上。
	// 并且获取当前用户是否加入某个房间，如果加入了，要直接跳转到指定房间内部。
	c.JSON(http.StatusOK, nil)
}

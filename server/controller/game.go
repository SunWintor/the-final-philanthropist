package controller

import (
	"github.com/SunWintor/tfp/server/controller/filter"
	"github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func handleGame(r *gin.Engine) {
	home := r.Group("/game", filter.JWTAuthHandler())

	home.GET("/info", gameInfo)
	home.POST("/donated", donated)
}

func gameInfo(c *gin.Context) {
	var r *model.GameInfoReq
	if err := c.BindQuery(&r); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	info, err := svr.GameInfo(c, r)
	gin_util.AutoResult(c, info, err)
}

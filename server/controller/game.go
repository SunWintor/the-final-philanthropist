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
	var g *model.GameInfoReq
	if err := c.BindQuery(&g); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	info, err := svr.GameInfo(c, g)
	gin_util.AutoResult(c, info, err)
}

func donated(c *gin.Context) {
	var d *model.DonatedReq
	if err := c.ShouldBindJSON(&d); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	err := svr.Donated(c, d)
	gin_util.AutoResult(c, nil, err)
}

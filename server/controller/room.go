package controller

import (
	"github.com/SunWintor/tfp/server/controller/filter"
	"github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func handleRoom(r *gin.Engine) {
	room := r.Group("/room", filter.JWTAuthHandler())

	room.POST("/join/random", joinRandomRoom, filter.EqualUltraViresHandler()) // 随机进入一个房间
	room.GET("/info", homeInfo)                                                // 获取某个房间的相关信息，包括房间里有谁，自己在不在这个房间里啥的
	room.POST("/ready", homeInfo, filter.EqualUltraViresHandler())             // 在房间内准备开始游戏，所有人都准备了游戏将会开始
	room.POST("/ready/cancel", homeInfo, filter.EqualUltraViresHandler())      // 取消准备
}

func joinRandomRoom(c *gin.Context) {
	var j *model.JoinRandomRoomReq
	if err := c.ShouldBindJSON(&j); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	roomId, err := svr.JoinRandomRoom(c, j)
	if err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	gin_util.SuccessWithObject(c, roomId)
}

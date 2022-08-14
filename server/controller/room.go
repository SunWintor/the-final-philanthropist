package controller

import (
	"github.com/SunWintor/tfp/server/controller/filter"
	"github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func handleRoom(r *gin.Engine) {
	room := r.Group("/room") //, filter.JWTAuthHandler())

	room.POST("/join/random", joinRandomRoom, filter.EqualUltraViresHandler())
	room.POST("/exit", exitRoom, filter.EqualUltraViresHandler())

	room.POST("/ready", ready, filter.EqualUltraViresHandler())
	room.POST("/ready/cancel", readyCancel, filter.EqualUltraViresHandler())

	room.GET("/info", roomInfo)
}

func ready(c *gin.Context) {
	var u *model.UserIdReq
	if err := c.ShouldBindJSON(&u); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	err := svr.Ready(c, u)
	gin_util.AutoResult(c, nil, err)
}

func readyCancel(c *gin.Context) {
	var u *model.UserIdReq
	if err := c.ShouldBindJSON(&u); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	err := svr.ReadyCancel(c, u)
	gin_util.AutoResult(c, nil, err)
}

func joinRandomRoom(c *gin.Context) {
	var u *model.UserIdReq
	if err := c.ShouldBindJSON(&u); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	roomId, err := svr.JoinRandomRoom(c, u)
	gin_util.AutoResult(c, roomId, err)
}

func exitRoom(c *gin.Context) {
	var u *model.UserIdReq
	if err := c.ShouldBindJSON(&u); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	err := svr.ExitRoom(c, u)
	gin_util.AutoResult(c, nil, err)
}

func roomInfo(c *gin.Context) {
	var r *model.RoomInfoReq
	if err := c.BindQuery(&r); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	roomInfo, err := svr.RoomInfo(c, r)
	gin_util.AutoResult(c, roomInfo, err)
}

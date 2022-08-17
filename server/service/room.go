package service

import (
	"github.com/SunWintor/tfp/server/core/room"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) Ready(c *gin.Context, arg *model.UserIdReq) (err error) {
	var r *room.Room
	if r, err = changeUserReadyStatus(arg.UserId, true); err != nil {
		return
	}
	// 在玩家点击准备按钮的时候，游戏开始失败不需要通知玩家。
	r.GameStart()
	return
}

func (s *Service) ReadyCancel(c *gin.Context, arg *model.UserIdReq) (err error) {
	_, err = changeUserReadyStatus(arg.UserId, false)
	return
}

func changeUserReadyStatus(userId int64, ready bool) (r *room.Room, err error) {
	if r, err = getUserStatusReadyRoom(userId); err != nil {
		return
	}
	if err = r.Ready(userId, ready); err != nil {
		return
	}
	return
}

func getUserStatusReadyRoom(userId int64) (r *room.Room, err error) {
	r, err = room.GetUserRoom(userId)
	if err != nil {
		return
	}
	if r.Status != room.GameReady {
		err = ecode.RoomNotReadyError
		return
	}
	return
}

func (s *Service) JoinRandomRoom(c *gin.Context, arg *model.UserIdReq) (res *model.JoinRandomRoomReply, err error) {
	roomId := room.GetUserRoomId(arg.UserId)
	if roomId != "" {
		res.RoomId = roomId
		return
	}
	r := room.GetJoinableRoom()
	res.RoomId = r.RoomId
	roomUser := room.GenerateRoomUser(arg.UserId)
	err = r.Join(roomUser)
	return
}

func (s *Service) ExitRoom(c *gin.Context, arg *model.UserIdReq) (err error) {
	var r *room.Room
	r, err = room.GetUserRoom(arg.UserId)
	if err != nil {
		return
	}
	err = r.Exit(arg.UserId)
	return
}

func (s *Service) RoomInfo(c *gin.Context, arg *model.RoomInfoReq) (res *model.RoomInfoReply, err error) {
	r := room.GetRoom(arg.RoomId)
	if r == nil {
		err = ecode.RoomNotExistsError
		return
	}
	res = r.ToReply()
	return
}

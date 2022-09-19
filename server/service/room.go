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
	var endChan <-chan struct{}
	if endChan, err = r.GameStart(); err != nil {
		// 在玩家点击准备按钮的时候，游戏开始失败不需要通知玩家。
		err = nil
		return
	}
	go func() {
		<-endChan
		s.GameEnd(&gin.Context{}, r)
	}()
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
	r, err = room.UserRoom(userId)
	if err != nil {
		return
	}
	if r.Status != room.GameReady {
		err = ecode.RoomNotReadyError
		return
	}
	return
}

func (s *Service) JoinRandomRoom(c *gin.Context, arg *model.UserIdReq) (res *model.RoomInfoReply, err error) {
	res = new(model.RoomInfoReply)
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if r != nil {
		res = r.ToReply(arg.UserId)
		err = ecode.AlreadyInRoomError
		return
	}
	return s.joinRoom(c, arg.UserId, room.GetJoinableRoom())
}

func (s *Service) JoinNew(c *gin.Context, arg *model.UserIdReq) (res *model.RoomInfoReply, err error) {
	res = new(model.RoomInfoReply)
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if r != nil {
		res = r.ToReply(arg.UserId)
		err = ecode.AlreadyInRoomError
		return
	}
	return s.joinRoom(c, arg.UserId, room.CreateEmptyRoom())
}

func (s *Service) JoinRoomId(c *gin.Context, arg *model.JoinRoomIdReq) (res *model.RoomInfoReply, err error) {
	res = new(model.RoomInfoReply)
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if r != nil {
		res = r.ToReply(arg.UserId)
		err = ecode.AlreadyInRoomError
		return
	}
	r = room.GetRoom(arg.RoomId)
	if r == nil {
		err = ecode.RoomNotExistsError
		return
	}
	if r.Status != room.GameReady {
		err = ecode.RoomNotReadyError
		return
	}
	return s.joinRoom(c, arg.UserId, r)
}

func (s *Service) joinRoom(c *gin.Context, userId int64, r *room.Room) (res *model.RoomInfoReply, err error) {
	var rankInfo *model.TfpUserRank
	if rankInfo, err = s.GetUserRankInfo(c, userId); err != nil {
		return
	}
	roomUser := room.GenerateRoomUser(userId, rankInfo.Ranking)
	err = r.Join(roomUser)
	return r.ToReply(userId), err
}

func (s *Service) ExitRoom(c *gin.Context, arg *model.UserIdReq) (err error) {
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if err != nil {
		return
	}
	if r == nil {
		err = ecode.RoomNotExistsError
		return
	}
	if r.Status == room.Gaming {
		err = ecode.GamingError
		return
	}
	err = r.Exit(arg.UserId)
	return
}

func (s *Service) RoomInfo(c *gin.Context, arg *model.UserIdReq) (res *model.RoomInfoReply, err error) {
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if err != nil {
		return
	}
	if r == nil {
		err = ecode.RoomNotExistsError
		return
	}
	res = r.ToReply(arg.UserId)
	return
}

package service

import (
	"github.com/SunWintor/tfp/server/core"
	"github.com/SunWintor/tfp/server/core/room"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) JoinRandomRoom(c *gin.Context, arg *model.JoinRandomRoomReq) (res *model.JoinRandomRoomReply, err error) {
	roomId := room.GetCurrentUserRoomId(arg.UserId)
	if roomId != "" {
		res.RoomId = roomId
		return
	}
	r := room.GetJoinableRoom()
	player := core.GeneratePlayer(arg.UserId)
	res.RoomId = r.RoomId
	err = r.Join(player)
	return
}

package service

import (
	"github.com/SunWintor/tfp/server/core"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) JoinRandomRoom(c *gin.Context, arg *model.JoinRandomRoomReq) (roomId string, err error) {
	roomId = core.GetRoomIdByUserId(arg.UserID)
	if roomId == "" {
		room := core.GetJoinableRoom()
		roomId = room.RoomID
	}
	return
}

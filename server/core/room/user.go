package room

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/core/user"
	"github.com/SunWintor/tfp/server/model"
)

type RoomUser struct {
	UserId   int64
	Username string
	RoomId   string
	IsReady  bool
}

func GenerateRoomUser(userId int64) *RoomUser {
	userName := ""
	u, ok := user.GetCopyById(userId)
	if ok {
		userName = u.Username
	}
	return &RoomUser{
		UserId:   userId,
		Username: userName,
	}
}

func (p *RoomUser) ToReply() *model.RoomUser {
	return &model.RoomUser{
		UserId:   p.UserId,
		IsReady:  p.IsReady,
		Username: p.Username,
	}
}

func (p *RoomUser) ToPlayer() *game.Player {
	return &game.Player{
		PlayerId: common.GetRandomPlayerId(),
		UserId:   p.UserId,
		Username: p.Username,
	}
}

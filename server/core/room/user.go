package room

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/core/user"
	"github.com/SunWintor/tfp/server/model"
)

type RoomUser struct {
	UserId  int64
	RoomId  string
	IsReady bool
}

func GenerateRoomUser(userId int64) *RoomUser {
	return &RoomUser{
		UserId: userId,
	}
}

func (p *RoomUser) ToReply() *model.RoomUser {
	return &model.RoomUser{
		UserId:  p.UserId,
		IsReady: p.IsReady,
	}
}

func (p *RoomUser) ToPlayer() *game.Player {
	userName := ""
	u, ok := user.GetCopyById(p.UserId)
	if ok {
		userName = u.Username
	}
	return &game.Player{
		PlayerId: common.GetRandomPlayerId(),
		UserId:   p.UserId,
		Username: userName,
	}
}

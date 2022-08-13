package core

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/core/game/hero"
	"github.com/SunWintor/tfp/server/model"
)

type Player struct {
	PlayerId string
	UserId   int64
	RoomId   string
	IsReady  bool
	Hero     *hero.Hero
}

func GeneratePlayer(userId int64) *Player {
	return &Player{
		PlayerId: common.GetRandomPlayerId(),
		UserId:   userId,
	}
}

func (p *Player) ToReply() *model.Player {
	return &model.Player{
		PlayerId: p.PlayerId,
		UserId:   p.UserId,
		IsReady:  p.IsReady,
		Hero:     (*p.Hero).ToReply(),
	}
}

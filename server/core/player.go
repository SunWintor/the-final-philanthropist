package core

import "github.com/SunWintor/tfp/server/core/game/hero"

type Player struct {
	PlayerID int64
	UserID   int64
	RoomID   string
	IsReady  bool
	Hero     *hero.Hero
}

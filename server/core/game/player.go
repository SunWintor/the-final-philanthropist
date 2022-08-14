package game

import "github.com/SunWintor/tfp/server/core/game/hero"

type Player struct {
	UserId   int64
	PlayerId string
	UserName string
	Hero     hero.Hero
}

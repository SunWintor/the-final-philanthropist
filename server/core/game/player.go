package game

import "github.com/SunWintor/tfp/server/core/game/hero"

type Player struct {
	UserId        int64
	PlayerId      string
	Username      string
	Hero          hero.Hero
	BankruptRound int64
	RoomRank      int64
	Ranking       float64
	RankingUp     float64
}

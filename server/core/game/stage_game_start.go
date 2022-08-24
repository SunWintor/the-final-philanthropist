package game

import (
	"github.com/SunWintor/tfp/server/core/game/hero"
	"time"
)

type gameStart struct {
	baseStage
}

func (g *gameStart) templateInit() {
	g.stage = GameStartStage
	g.stageName = "游戏开始阶段"
	g.duration = 3 * time.Second
}

func (g *gameStart) Run(ctx *ProcessContext) <-chan time.Time {
	for _, player := range ctx.PlayerMap {
		h := &hero.Foo2Die{}
		h.Init()
		player.Hero = h
	}
	return g.baseStage.Run(ctx)
}

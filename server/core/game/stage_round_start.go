package game

import (
	"time"
)

type roundStart struct {
	baseStage
}

func (g *roundStart) templateInit() {
	g.stage = RoundStartStage
	g.stageName = "回合开始阶段"
	g.duration = 3 * time.Second
}

func (g *roundStart) Run(ctx *ProcessContext) <-chan time.Time {
	ctx.Round++
	ctx.initCurrentRound()
	return g.baseStage.Run(ctx)
}

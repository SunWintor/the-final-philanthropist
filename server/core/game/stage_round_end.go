package game

import (
	"time"
)

type roundEnd struct {
	baseStage
}

func (g *roundEnd) templateInit() {
	g.stage = RoundEndStage
	g.stageName = "回合结束阶段"
	g.duration = 3 * time.Second
}

func (g *roundEnd) Run(ctx *ProcessContext) <-chan time.Time {
	ctx.roundToHistory()
	ctx.judgementGameEnd()
	return g.baseStage.Run(ctx)
}

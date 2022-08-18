package stage

import (
	"github.com/SunWintor/tfp/server/core/game/process"
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

func (g *roundEnd) Run(ctx *process.ProcessContext) <-chan time.Time {
	ctx.RoundToHistory()
	ctx.JudgementGameEnd()
	return g.baseStage.Run(ctx)
}

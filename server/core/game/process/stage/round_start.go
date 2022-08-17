package stage

import (
	"github.com/SunWintor/tfp/server/core/game/process"
	"time"
)

type roundStart struct {
	baseStage
}

func (g *roundStart) templateInit() {
	g.Stage = RoundStartStage
	g.StageName = "回合开始阶段"
	g.DurationSecond = 3 * time.Second
}

func (g *roundStart) Run(ctx *process.ProcessContext) <-chan time.Time {
	ctx.Round++
	ctx.InitCurrentRound()
	return g.baseStage.Run(ctx)
}

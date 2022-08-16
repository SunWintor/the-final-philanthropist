package stage

import (
	"github.com/SunWintor/tfp/server/core/game/process"
	"time"
)

type publicOpinion struct {
	baseStage
}

func (g *publicOpinion) templateInit() {
	g.Stage = PublicOpinionStage
	g.StageName = "舆论惩罚展示阶段"
	g.DurationSecond = 5 * time.Second
}

func (g *publicOpinion) Run(ctx *process.ProcessContext) <-chan time.Time {
	return g.baseStage.Run(ctx)
}

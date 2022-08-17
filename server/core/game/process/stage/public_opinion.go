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
	g.StageName = "舆论惩罚阶段"
	g.DurationSecond = 5 * time.Second
}

func (g *publicOpinion) Run(ctx *process.ProcessContext) <-chan time.Time {
	currentRound := ctx.CurrentRoundInfo
	minPlayerList := currentRound.DonatedPlayer(currentRound.MinDonated())
	currentRound.PunishmentPlayer(minPlayerList)
	ctx.SyncPunishment()
	return g.baseStage.Run(ctx)
}

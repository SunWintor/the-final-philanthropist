package game

import (
	"time"
)

type publicOpinion struct {
	baseStage
}

func (g *publicOpinion) templateInit() {
	g.stage = PublicOpinionStage
	g.stageName = "舆论惩罚阶段"
	g.duration = 5 * time.Second
}

func (g *publicOpinion) Run(ctx *ProcessContext) <-chan time.Time {
	currentRound := ctx.CurrentRoundInfo
	minPlayerList := currentRound.PlayerByDonated(currentRound.MinDonated())
	currentRound.PunishmentGen(minPlayerList)
	ctx.SyncPunishment()
	return g.baseStage.Run(ctx)
}

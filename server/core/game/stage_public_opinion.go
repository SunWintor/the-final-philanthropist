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
	currentRound.reckonPunishment()
	ctx.decPlayerPunishmentMoney()
	ctx.judgementGameEnd()
	return g.baseStage.Run(ctx)
}

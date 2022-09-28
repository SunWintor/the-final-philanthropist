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
	for _, v := range currentRound.DonatedInfoList {
		v.PunishmentMoney = ctx.PlayerMap[v.PlayerId].Hero.OnPublicOpinion(v.PunishmentMoney)
	}
	ctx.decPlayerPunishmentMoney()
	ctx.judgementGameEnd()
	return g.baseStage.Run(ctx)
}

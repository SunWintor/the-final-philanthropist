package game

import "time"

type bankruptcy struct {
	baseStage
}

func (g *bankruptcy) templateInit() {
	g.stage = BankruptcyStage
	g.stageName = "破产判定阶段"
	g.duration = 5 * time.Second
}

// 未来，可能会有英雄技能让角色在此起死回生。

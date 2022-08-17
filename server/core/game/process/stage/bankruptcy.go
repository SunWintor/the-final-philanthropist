package stage

import "time"

type bankruptcy struct {
	baseStage
}

func (g *bankruptcy) templateInit() {
	g.Stage = BankruptcyStage
	g.StageName = "破产判定阶段"
	g.DurationSecond = 5 * time.Second
}

// 未来，可能会有英雄技能让角色在此起死回生。

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

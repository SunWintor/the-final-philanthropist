package game

import (
	"time"
)

type donated struct {
	baseStage
}

func (g *donated) templateInit() {
	g.stage = DonatedStage
	g.stageName = "捐赠阶段"
	g.duration = 20 * time.Second
}

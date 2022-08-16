package stage

import (
	"time"
)

type donated struct {
	baseStage
}

func (g *donated) templateInit() {
	g.Stage = DonatedStage
	g.StageName = "捐赠阶段"
	g.DurationSecond = 20 * time.Second
}

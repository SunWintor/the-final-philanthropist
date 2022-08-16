package stage

import "time"

type roundEnd struct {
	baseStage
}

func (g *roundEnd) templateInit() {
	g.Stage = RoundEndStage
	g.StageName = "回合结束阶段"
	g.DurationSecond = 3 * time.Second
}

package stage

import "time"

// 这个阶段是为了防止捐赠阶段末尾来的请求污染捐赠结果
type donatedEnd struct {
	baseStage
}

func (g *donatedEnd) templateInit() {
	g.stage = DonatedEndStage
	g.stageName = "捐赠结束阶段"
	g.duration = 3 * time.Second
}

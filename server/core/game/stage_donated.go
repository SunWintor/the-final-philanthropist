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
	g.duration = 18 * time.Second
}

func (s *donated) Run(ctx *ProcessContext) <-chan time.Time {
	return time.After(time.Duration(s.GetDurationSecond(ctx)) * time.Second)
}

func (s *donated) GetDurationSecond(ctx *ProcessContext) int64 {
	duration := s.duration + time.Duration(ctx.RichPlayerCount+ctx.Round)*time.Second
	return int64(duration / time.Second)
}

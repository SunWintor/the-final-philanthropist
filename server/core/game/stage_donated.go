package game

import (
	"math"
	"time"
)

type donated struct {
	baseStage
}

func (g *donated) templateInit() {
	g.stage = DonatedStage
	g.stageName = "捐赠阶段"
	g.duration = 60 * time.Second
}

func (s *donated) Run(ctx *ProcessContext) <-chan time.Time {
	return time.After(time.Duration(s.GetDurationSecond(ctx)) * time.Second)
}

func (s *donated) GetDurationSecond(ctx *ProcessContext) int64 {
	coefficient := ctx.RichPlayerCount*2 - int64(math.Abs(float64(ctx.Round-7)))*7
	duration := s.duration + time.Duration(coefficient)*time.Second
	if duration < 20*time.Second {
		duration = 20 * time.Second
	}
	return int64(duration / time.Second)
}

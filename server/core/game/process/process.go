package process

import (
	"github.com/SunWintor/tfp/server/core/game/process/stage"
	"github.com/SunWintor/tfp/server/model"
	"time"
)

type Process struct {
	Stage          stage.Stage
	StageStartTime time.Time
	ProcessContext *ProcessContext
}

func (p *Process) ToGameInfoReply() *model.GameInfo {
	return &model.GameInfo{}
}

func (p *Process) Start() {
	p.Stage = stage.StageMap[stage.GameStartStage]
	for {
		select {
		case <-p.ProcessContext.EndGame:
			return
		default:
			p.StageStartTime = time.Now()
			<-p.Stage.Run(p.ProcessContext)
			p.Stage = p.Stage.Next(p.ProcessContext)
		}
	}
}

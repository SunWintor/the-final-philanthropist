package process

import (
	"github.com/SunWintor/tfp/server/core/game/process/stage"
)

type Process struct {
	Stage          stage.Stage
	ProcessContext *ProcessContext
}

func (p *Process) Start() {
	p.Stage = stage.StageMap[stage.GameStartStage]
	for {
		select {
		case <-p.ProcessContext.EndGame:
			return
		default:
			<-p.Stage.Run(p.ProcessContext)
			p.Stage = p.Stage.Next(p.ProcessContext)
		}
	}
}

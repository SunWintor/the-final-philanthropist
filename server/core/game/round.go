package game

import "github.com/SunWintor/tfp/server/core/game/stage"

type Round struct {
	RoundNo int64
	Stage   stage.Stage
}

func (r *Round) Next() {
	r.RoundNo++
	r.Stage.Run()
}

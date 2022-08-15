package process

import "time"

type Round struct {
	RoundNo int64
	Stage   *Stage
}

func (r *Round) NextStage() <-chan time.Time {
	r.Stage = r.Stage.Next()
	return r.Stage.Run()
}

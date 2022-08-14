package stage

import "time"

type PublicOpinionStage struct {
	baseStage
}

func (r *PublicOpinionStage) Run() (hasNext bool) {
	<-time.After(r.DurationSecond)
	return true
}

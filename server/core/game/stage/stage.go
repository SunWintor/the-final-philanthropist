package stage

import (
	"time"
)

type Stage interface {
	Run() (hasNext bool)
}

type baseStage struct {
	Stage          int64
	StageName      string
	DurationSecond time.Duration
	Next           Stage
}

func (b *baseStage) Run() (hasNext bool) {
	<-time.After(b.DurationSecond)
	return true
}

var Donated *DonatedStage
var PublicOpinion *PublicOpinionStage

func init() {
	Donated = &DonatedStage{}
	Donated.Stage = 1
	Donated.StageName = "捐赠阶段"
	Donated.DurationSecond = 20 * time.Second

	PublicOpinion = &PublicOpinionStage{}
	PublicOpinion.Stage = 2
	PublicOpinion.StageName = "舆论惩罚展示阶段"
	PublicOpinion.DurationSecond = 5 * time.Second

	Donated.Next = PublicOpinion
	PublicOpinion.Next = Donated
}

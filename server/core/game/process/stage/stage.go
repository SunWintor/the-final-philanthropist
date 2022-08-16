package stage

import (
	"github.com/SunWintor/tfp/server/core/game/process"
	"time"
)

type Stage interface {
	Run(ctx *process.ProcessContext) <-chan time.Time
	Next(ctx *process.ProcessContext) Stage
	setNext(stage Stage)
	templateInit()
}

type baseStage struct {
	Stage          int64
	StageName      string
	DurationSecond time.Duration
	next           Stage
}

func (s *baseStage) Run(ctx *process.ProcessContext) <-chan time.Time {
	return time.After(s.DurationSecond)
}

func (s *baseStage) Next(ctx *process.ProcessContext) Stage {
	return s.next
}

func (s *baseStage) setNext(stage Stage) {
	s.next = stage
}

var StageMap map[int64]Stage

const (
	GameStartStage = iota
	RoundStartStage
	DonatedStage
	PublicOpinionStage
	BankruptcyStage
	RoundEndStage = 99
)

func init() {
	StageMap = make(map[int64]Stage)
	StageMap[GameStartStage] = &gameStart{}
	StageMap[RoundStartStage] = &roundStart{}
	StageMap[DonatedStage] = &donated{}
	StageMap[PublicOpinionStage] = &publicOpinion{}
	StageMap[BankruptcyStage] = &bankruptcy{}
	StageMap[RoundEndStage] = &roundEnd{}

	for _, v := range StageMap {
		v.templateInit()
	}

	StageMap[GameStartStage].setNext(StageMap[RoundStartStage])
	StageMap[RoundStartStage].setNext(StageMap[DonatedStage])
	StageMap[DonatedStage].setNext(StageMap[PublicOpinionStage])
	StageMap[PublicOpinionStage].setNext(StageMap[BankruptcyStage])
	StageMap[BankruptcyStage].setNext(StageMap[RoundEndStage])
	StageMap[RoundEndStage].setNext(StageMap[DonatedStage])
}

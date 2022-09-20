package game

import (
	"time"
)

type Stage interface {
	Run(ctx *ProcessContext) <-chan time.Time
	Next(ctx *ProcessContext) Stage
	setNext(stage Stage)
	templateInit()
	GetStage() int64
	GetName() string
	GetDurationSecond(ctx *ProcessContext) int64
	runHeroSkill(ctx *ProcessContext)
}

type baseStage struct {
	stage     int64
	stageName string
	duration  time.Duration
	next      Stage
}

func (s *baseStage) GetStage() int64 {
	return s.stage
}

func (s *baseStage) GetName() string {
	return s.stageName
}

func (s *baseStage) GetDurationSecond(ctx *ProcessContext) int64 {
	return int64(s.duration / time.Second)
}

func (s *baseStage) Run(ctx *ProcessContext) <-chan time.Time {
	return time.After(s.duration)
}

func (s *baseStage) Next(ctx *ProcessContext) Stage {
	return s.next
}

func (s *baseStage) runHeroSkill(ctx *ProcessContext) {
	return
}

func (s *baseStage) setNext(stage Stage) {
	s.next = stage
}

var StageMap map[int64]Stage

const (
	GameStartStage = iota
	RoundStartStage
	DonatedStage
	DonatedEndStage
	PublicOpinionStage
	BankruptcyStage
	RoundEndStage = 99
)

func init() {
	StageMap = make(map[int64]Stage)
	StageMap[GameStartStage] = &gameStart{}
	StageMap[RoundStartStage] = &roundStart{}
	StageMap[DonatedStage] = &donated{}
	StageMap[DonatedEndStage] = &donatedEnd{}
	StageMap[PublicOpinionStage] = &publicOpinion{}
	StageMap[BankruptcyStage] = &bankruptcy{}
	StageMap[RoundEndStage] = &roundEnd{}

	for _, v := range StageMap {
		v.templateInit()
	}

	// todo 可以优化为生成的时候就存一个nextId，然后需要的时候去map取
	StageMap[GameStartStage].setNext(StageMap[RoundStartStage])
	StageMap[RoundStartStage].setNext(StageMap[DonatedStage])
	StageMap[DonatedStage].setNext(StageMap[DonatedEndStage])
	StageMap[DonatedEndStage].setNext(StageMap[PublicOpinionStage])
	StageMap[PublicOpinionStage].setNext(StageMap[RoundEndStage])
	StageMap[RoundEndStage].setNext(StageMap[RoundStartStage])
}

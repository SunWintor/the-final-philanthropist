package process

import (
	"time"
)

type Stage struct {
	Stage          int64
	StageName      string
	DurationSecond time.Duration
	next           *Stage
}

func (s *Stage) Run() <-chan time.Time {
	return time.After(s.DurationSecond)
}

func (s *Stage) Next() *Stage {
	return s.next
}

var StageMap map[int64]*Stage

const (
	GameStartStage = iota
	RoundStartStage
	DonatedStage
	PublicOpinionStage
	BankruptcyStage
	RoundEndStage = 99
)

func init() {
	StageMap = make(map[int64]*Stage)
	StageMap[GameStartStage] = &Stage{Stage: 0, StageName: "游戏开始阶段", DurationSecond: 3 * time.Second}
	StageMap[RoundStartStage] = &Stage{Stage: 1, StageName: "回合开始阶段", DurationSecond: 3 * time.Second}
	StageMap[DonatedStage] = &Stage{Stage: 2, StageName: "捐赠阶段", DurationSecond: 20 * time.Second}
	StageMap[PublicOpinionStage] = &Stage{Stage: 3, StageName: "舆论惩罚展示阶段", DurationSecond: 5 * time.Second}
	StageMap[BankruptcyStage] = &Stage{Stage: 4, StageName: "破产判定阶段", DurationSecond: 5 * time.Second}
	StageMap[RoundEndStage] = &Stage{Stage: 99, StageName: "回合结束阶段", DurationSecond: 3 * time.Second}

	StageMap[GameStartStage].next = StageMap[RoundStartStage]
	StageMap[RoundStartStage].next = StageMap[DonatedStage]
	StageMap[DonatedStage].next = StageMap[PublicOpinionStage]
	StageMap[PublicOpinionStage].next = StageMap[BankruptcyStage]
	StageMap[BankruptcyStage].next = StageMap[RoundEndStage]
	StageMap[RoundEndStage].next = StageMap[DonatedStage]
}

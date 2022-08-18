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

func (p *Process) ToGameInfoReply(userId int64) *model.GameInfo {
	return &model.GameInfo{
		RoundInfo:        p.toRoundInfoReply(),
		PlayerGameInfo:   p.toUserInfoReply(userId),
		RoundHistoryList: p.toRoundHistoryReplyList(),
	}
}

func (p *Process) toRoundHistoryReplyList() []*model.RoundHistory {
	var res []*model.RoundHistory
	for _, roundHistory := range p.ProcessContext.RoundHistory {
		res = append(res, roundHistory.ToReply())
	}
	return res
}

func (p *Process) toUserInfoReply(userId int64) *model.PlayerGameInfo {
	userInfo := &model.PlayerGameInfo{}
	for playerId, player := range p.ProcessContext.PlayerMap {
		if player.UserId == userId {
			userInfo.PlayerId = playerId
			userInfo.Username = player.Username
			userInfo.HeroInfo = player.Hero.ToReply()
		}
	}
	return userInfo
}

func (p *Process) toRoundInfoReply() *model.RoundInfo {
	return &model.RoundInfo{
		RoundNo:       p.ProcessContext.Round,
		PublicOpinion: p.ProcessContext.CurrentRoundInfo.PublicOpinion,
		Stage: &model.Stage{
			Stage:          p.Stage.GetStage(),
			StartTime:      p.StageStartTime.UnixNano() / 1000 / 1000,
			Name:           p.Stage.GetName(),
			DurationSecond: p.Stage.GetDurationSecond(),
		},
	}
}

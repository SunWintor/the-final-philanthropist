package game

import (
	"github.com/SunWintor/tfp/server/model"
	"time"
)

type Process struct {
	Stage          Stage
	StageStartTime time.Time
	ProcessContext *ProcessContext
}

func (p *Process) Start() {
	p.Stage = StageMap[GameStartStage]
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
	playerGameInfo := p.toUserInfoReply(userId)
	return &model.GameInfo{
		RoundInfo:        p.toRoundInfoReply(playerGameInfo.PlayerId),
		PlayerGameInfo:   playerGameInfo,
		RoundHistoryList: p.toRoundHistoryReplyList(),
	}
}

func (p *Process) emptyRoundHistoryReply() *model.RoundHistory {
	var roundPlayerInfo []*model.RoundDonatedInfo
	for _, player := range p.ProcessContext.PlayerMap {
		heroName := ""
		limitMoney := int64(0)
		if player.Hero != nil {
			heroName = player.Hero.GetName()
			limitMoney = player.Hero.GetMoneyLimit()
		}
		empty := &model.RoundDonatedInfo{
			PlayerId:        player.PlayerId,
			Username:        player.Username,
			HeroName:        heroName,
			CurrentMoney:    limitMoney, // 这里没有赋值错误，就是要展示玩家的初始值，因为这个empty是展示在图标首部的
			DonatedMoney:    0,
			PunishmentMoney: 0,
			Bankrupt:        false,
		}
		roundPlayerInfo = append(roundPlayerInfo, empty)
	}
	return &model.RoundHistory{
		RoundNo:              0,
		RoundDonatedInfoList: roundPlayerInfo,
	}
}

func (p *Process) toRoundHistoryReplyList() []*model.RoundHistory {
	var res []*model.RoundHistory
	// 前端展示的图表的第一行数据
	res = append(res, p.emptyRoundHistoryReply())
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

func (p *Process) toRoundInfoReply(playerId string) *model.RoundInfo {
	currentRoundInfo := &model.RoundHistory{
		RoundNo: p.ProcessContext.Round,
	}
	var roundInfoList []*model.RoundDonatedInfo
	for _, donatedInfo := range p.ProcessContext.CurrentRoundInfo.DonatedInfoList {
		donatedInfoReply := donatedInfo.ToReply()
		if p.Stage.GetStage() == DonatedStage && donatedInfo.PlayerId != playerId {
			donatedInfoReply.DonatedMoney = -1
		}
		roundInfoList = append(roundInfoList, donatedInfoReply)
	}
	currentRoundInfo.RoundDonatedInfoList = roundInfoList
	return &model.RoundInfo{
		RoundNo:       p.ProcessContext.Round,
		PublicOpinion: p.ProcessContext.CurrentRoundInfo.PublicOpinion,
		Stage: &model.Stage{
			Stage:          p.Stage.GetStage(),
			StartTime:      p.StageStartTime.UnixNano() / 1000 / 1000,
			Name:           p.Stage.GetName(),
			DurationSecond: p.Stage.GetDurationSecond(),
		},
		CurrentRoundInfo: currentRoundInfo,
	}
}

package game

import (
	"github.com/SunWintor/tfp/server/model"
	"log"
	"sort"
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

func (p *Process) emptyRoundHistoryReplyMap() (res map[string]*model.RoundDonatedInfo) {
	res = make(map[string]*model.RoundDonatedInfo, len(p.ProcessContext.PlayerMap)*2)
	for _, player := range p.ProcessContext.PlayerMap {
		heroName := ""
		limitMoney := int64(0)
		if player.Hero != nil {
			heroName = player.Hero.GetName()
			limitMoney = player.Hero.GetMoneyLimit()
		}
		empty := &model.RoundDonatedInfo{
			PlayerId:            player.PlayerId,
			Username:            player.Username,
			HeroName:            heroName,
			CurrentMoneyList:    []int64{limitMoney}, // 这里没有赋值错误，就是要展示玩家的初始值，因为这个empty是展示在图标首部的
			DonatedMoneyList:    []int64{0},
			PunishmentMoneyList: []int64{0},
			BankruptList:        []bool{false},
		}
		res[player.PlayerId] = empty
	}
	return
}

func (p *Process) toRoundHistoryReplyList() (res []*model.RoundDonatedInfo) {
	replyMap := p.emptyRoundHistoryReplyMap()
	for _, roundHistory := range p.ProcessContext.RoundHistory {
		for _, donatedInfo := range roundHistory.DonatedInfoList {
			info, ok := replyMap[donatedInfo.PlayerId]
			if !ok {
				log.Printf("toRoundHistoryReplyList err info is nil %+v %+v", replyMap, p)
				continue
			}
			info.CurrentMoneyList = append(info.CurrentMoneyList, donatedInfo.CurrentMoney)
			info.DonatedMoneyList = append(info.DonatedMoneyList, donatedInfo.DonatedMoney)
			info.PunishmentMoneyList = append(info.PunishmentMoneyList, donatedInfo.PunishmentMoney)
			info.BankruptList = append(info.BankruptList, donatedInfo.Bankrupt)
		}
	}
	for _, v := range replyMap {
		res = append(res, v)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].PlayerId < res[j].PlayerId
	})
	return
}

func (p *Process) toUserInfoReply(userId int64) *model.PlayerGameInfo {
	userInfo := &model.PlayerGameInfo{}
	for playerId, player := range p.ProcessContext.PlayerMap {
		if player.UserId == userId {
			userInfo.PlayerId = playerId
			userInfo.Username = player.Username
			userInfo.HeroInfo = player.Hero.ToReply()
			userInfo.Ranking = player.Ranking
		}
	}
	return userInfo
}

func (p *Process) toRoundInfoReply(playerId string) *model.RoundInfo {
	return &model.RoundInfo{
		RoundNo:       p.ProcessContext.Round,
		PublicOpinion: p.ProcessContext.CurrentRoundInfo.PublicOpinion,
		Stage: &model.Stage{
			Stage:          p.Stage.GetStage(),
			StartTime:      p.StageStartTime.UnixNano() / 1000 / 1000,
			Name:           p.Stage.GetName(),
			DurationSecond: p.Stage.GetDurationSecond(p.ProcessContext),
		},
		PlayerInfoList: p.toPlayerInfoListReply(playerId),
	}
}

func (p *Process) toPlayerInfoListReply(playerId string) []*model.PlayerInfo {
	var playerInfoList []*model.PlayerInfo
	for _, donatedInfo := range p.ProcessContext.CurrentRoundInfo.DonatedInfoList {
		donatedInfoReply := donatedInfo.ToPlayerReply()
		if p.Stage.GetStage() == DonatedStage && donatedInfo.PlayerId != playerId {
			donatedInfoReply.DonatedMoney = -1
		}
		playerInfoList = append(playerInfoList, donatedInfoReply)
	}
	sort.Slice(playerInfoList, func(i, j int) bool {
		return playerInfoList[i].PlayerId < playerInfoList[j].PlayerId
	})
	return playerInfoList
}

package game

import (
	"fmt"
	"sort"
	"strconv"
)

type ProcessContext struct {
	Round            int64
	PlayerMap        map[string]*Player // 该map初始化后便不会变动，所以无需加锁便可操作。
	CurrentRoundInfo *RoundInfo
	RoundHistory     []*RoundInfo
	EndGame          chan struct{}
	RichPlayerCount  int64
}

const (
	PublicOpinionBet   = 2
	PublicOpinionLimit = 10
)

func reckonPublicOpinion(round int64) int64 {
	res := round + PublicOpinionBet
	if res > PublicOpinionLimit {
		res = PublicOpinionLimit
	}
	return res
}

func (p *ProcessContext) initCurrentRound() {
	p.CurrentRoundInfo = &RoundInfo{
		RoundNo:         p.Round,
		PublicOpinion:   reckonPublicOpinion(p.Round),
		DonatedInfoList: make([]*DonatedInfo, 0, len(p.PlayerMap)),
	}
	for playerId, player := range p.PlayerMap {
		donatedInfo := &DonatedInfo{
			PlayerId:        playerId,
			Username:        player.Username,
			Ranking:         player.Ranking,
			DonatedMoney:    p.CurrentRoundInfo.defaultDonatedMoney(), // 当玩家超时或者掉线的时候，捐赠额度为舆论惩罚+2，对局势影响会较小
			PunishmentMoney: 0,
		}
		donatedInfo.initHeroInfo(player.Hero)
		p.CurrentRoundInfo.DonatedInfoList = append(p.CurrentRoundInfo.DonatedInfoList, donatedInfo)
	}
}

func (p *ProcessContext) syncHeroMoneyToCurrentRound() {
	for _, d := range p.CurrentRoundInfo.DonatedInfoList {
		player, ok := p.PlayerMap[d.PlayerId]
		if !ok {
			continue
		}
		d.CurrentMoney = player.Hero.GetCurrentMoney()
	}
}

func (p *ProcessContext) decPlayerDonatedMoney() {
	for _, donatedInfo := range p.CurrentRoundInfo.DonatedInfoList {
		p.decPlayerMoney(donatedInfo.DonatedMoney, donatedInfo)
	}
}

func (p *ProcessContext) decPlayerPunishmentMoney() {
	for _, donatedInfo := range p.CurrentRoundInfo.DonatedInfoList {
		p.decPlayerMoney(donatedInfo.PunishmentMoney, donatedInfo)
	}
}

func (p *ProcessContext) decPlayerMoney(money int64, donatedInfo *DonatedInfo) {
	if money == 0 {
		return
	}
	if donatedInfo.Bankrupt {
		return
	}
	player, _ := p.PlayerMap[donatedInfo.PlayerId]
	player.Hero.DecMoney(money)
	if player.Hero.IsBankrupt() {
		donatedInfo.Bankrupt = true
	}
	p.syncHeroMoneyToCurrentRound()
	return
}

func (p *ProcessContext) rankSettlement() {
	var roundBankruptPlayer []*Player
	for _, player := range p.PlayerMap {
		if !player.Hero.IsBankrupt() {
			continue
		}
		if player.BankruptRound == 0 {
			player.BankruptRound = p.Round
			roundBankruptPlayer = append(roundBankruptPlayer, player)
		}
	}
	if len(roundBankruptPlayer) == 0 {
		return
	}
	sort.Slice(roundBankruptPlayer, func(i, j int) bool {
		return roundBankruptPlayer[i].Hero.GetCurrentMoney() < roundBankruptPlayer[j].Hero.GetCurrentMoney()
	})
	for _, player := range roundBankruptPlayer {
		player.RoomRank = p.RichPlayerCount
		p.RichPlayerCount--
	}
}

func (p *ProcessContext) roundToHistory() {
	p.RoundHistory = append(p.RoundHistory, p.CurrentRoundInfo)
}

func (p *ProcessContext) judgementGameEnd() {
	p.rankSettlement()
	if p.RichPlayerCount <= 1 {
		p.roundToHistory()
		p.endGameRankSettlement()
		close(p.EndGame)
	}
}

func (p *ProcessContext) endGameRankSettlement() {
	for _, player := range p.PlayerMap {
		if player.RoomRank == 0 {
			player.RoomRank = p.RichPlayerCount
			p.RichPlayerCount--
		}
	}
}

func (p *ProcessContext) GetRankingUp(userId int64) float64 {
	for _, v := range p.PlayerMap {
		if v.UserId == userId {
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v.RankingUp), 64)
			return value
		}
	}
	return 0
}

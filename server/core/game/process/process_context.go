package process

import (
	"github.com/SunWintor/tfp/server/core/game"
)

type ProcessContext struct {
	Round            int64
	PlayerMap        map[string]*game.Player // 该map初始化后便不会变动，所以无需加锁便可操作。
	CurrentRoundInfo *RoundInfo
	RoundHistory     []*RoundInfo
	EndGame          chan struct{}
}

func (p *ProcessContext) InitCurrentRound() {
	p.CurrentRoundInfo = &RoundInfo{
		RoundNo:         p.Round,
		PublicOpinion:   p.Round + 2,
		DonatedInfoList: make([]*DonatedInfo, len(p.PlayerMap)),
	}
	for playerId, player := range p.PlayerMap {
		p.CurrentRoundInfo.DonatedInfoList = append(p.CurrentRoundInfo.DonatedInfoList, &DonatedInfo{
			PlayerId:        playerId,
			CurrentMoney:    player.Hero.GetCurrentMoney(),
			MoneyLimit:      player.Hero.GetMoneyLimit(),
			DonatedMoney:    p.CurrentRoundInfo.PublicOpinion + 2, // 当玩家超时或者掉线的时候，捐赠额度为舆论惩罚+2，对局势影响会较小
			PunishmentMoney: 0,                                    // 显式声明，表示全部字段都处理过了
			Bankrupt:        false,                                // 同上
		})
	}
}

func (p *ProcessContext) SyncPunishment() {
	for _, donatedInfo := range p.CurrentRoundInfo.DonatedInfoList {
		if donatedInfo.PunishmentMoney == 0 {
			continue
		}
		player, _ := p.PlayerMap[donatedInfo.PlayerId]
		if player.Hero.IsBankrupt() {
			continue
		}
		player.Hero.Punishment(p.CurrentRoundInfo.PublicOpinion)
		if player.Hero.IsBankrupt() {
			donatedInfo.Bankrupt = true
		}
	}
}

func (p *ProcessContext) RoundToHistory() {
	p.RoundHistory = append(p.RoundHistory, p.CurrentRoundInfo)
}

func (p *ProcessContext) JudgementGameEnd() {
	lifePlayerCount := 0
	for _, player := range p.PlayerMap {
		if player.Hero.IsBankrupt() {
			continue
		}
		lifePlayerCount++
	}
	if lifePlayerCount <= 1 {
		close(p.EndGame)
	}
}

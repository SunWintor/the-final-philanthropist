package game

import (
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"math"
)

type RoundInfo struct {
	RoundNo         int64
	PublicOpinion   int64
	DonatedInfoList []*DonatedInfo
}

type DonatedInfo struct {
	PlayerId             string
	Username             string
	HeroName             string
	CurrentMoney         int64
	MoneyLimit           int64
	DonatedMoney         int64
	PunishmentMoney      int64
	Bankrupt             bool
	CurrentRoundBankrupt bool
}

func (r *RoundInfo) defaultDonatedMoney() int64 {
	return r.PublicOpinion + 2
}

func (g *RoundInfo) donated(playerId string, donated int64) error {
	for _, donatedInfo := range g.DonatedInfoList {
		if donatedInfo.PlayerId != playerId {
			continue
		}
		donatedInfo.DonatedMoney = donated
		if donatedInfo.CurrentMoney < donated {
			donatedInfo.DonatedMoney = donatedInfo.CurrentMoney
		}
		return nil
	}
	return ecode.PlayerNotExistsError
}

func (r *RoundInfo) minDonated() int64 {
	minDonated := int64(math.MaxInt64)
	for _, donatedInfo := range r.DonatedInfoList {
		if donatedInfo.DonatedMoney < minDonated {
			minDonated = donatedInfo.DonatedMoney
		}
	}
	return minDonated
}

func (r *RoundInfo) maxDonated() int64 {
	maxDonated := int64(0)
	for _, donatedInfo := range r.DonatedInfoList {
		if donatedInfo.DonatedMoney > maxDonated {
			maxDonated = donatedInfo.DonatedMoney
		}
	}
	return maxDonated
}

func (r *RoundInfo) playerByDonated(donated int64) []string {
	res := make([]string, 0)
	for _, donatedInfo := range r.DonatedInfoList {
		if donatedInfo.DonatedMoney == donated {
			res = append(res, donatedInfo.PlayerId)
		}
	}
	return res
}

func (r *RoundInfo) reckonPunishment() {
	minPlayerList := r.playerByDonated(r.minDonated())
	playerIdMap := make(map[string]struct{}, len(minPlayerList)*2)
	for _, playerId := range minPlayerList {
		playerIdMap[playerId] = struct{}{}
	}
	for _, donatedInfo := range r.DonatedInfoList {
		if _, ok := playerIdMap[donatedInfo.PlayerId]; ok {
			donatedInfo.PunishmentMoney = r.PublicOpinion
		}
	}
}

func (r *RoundInfo) ToReply() *model.RoundHistory {
	var roundPlayerInfo []*model.RoundDonatedInfo
	for _, donatedInfo := range r.DonatedInfoList {
		roundPlayerInfo = append(roundPlayerInfo, donatedInfo.ToReply())
	}
	return &model.RoundHistory{
		RoundNo:              r.RoundNo,
		RoundDonatedInfoList: roundPlayerInfo,
	}
}

func (d *DonatedInfo) ToReply() *model.RoundDonatedInfo {
	return &model.RoundDonatedInfo{
		PlayerId:        d.PlayerId,
		Username:        d.Username,
		HeroName:        d.HeroName,
		CurrentMoney:    d.CurrentMoney,
		DonatedMoney:    d.DonatedMoney,
		PunishmentMoney: d.PunishmentMoney,
		Bankrupt:        d.Bankrupt,
	}
}

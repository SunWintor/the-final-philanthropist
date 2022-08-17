package process

import "math"

type RoundInfo struct {
	RoundNo         int64
	PublicOpinion   int64
	DonatedInfoList []*DonatedInfo
}

type DonatedInfo struct {
	PlayerId        string
	CurrentMoney    int64
	MoneyLimit      int64
	DonatedMoney    int64
	PunishmentMoney int64
	Bankrupt        bool
}

func (r *RoundInfo) MinDonated() int64 {
	minDonated := int64(math.MaxInt64)
	for _, donatedInfo := range r.DonatedInfoList {
		if donatedInfo.DonatedMoney < minDonated {
			minDonated = donatedInfo.DonatedMoney
		}
	}
	return minDonated
}

func (r *RoundInfo) MaxDonated() int64 {
	maxDonated := int64(0)
	for _, donatedInfo := range r.DonatedInfoList {
		if donatedInfo.DonatedMoney > maxDonated {
			maxDonated = donatedInfo.DonatedMoney
		}
	}
	return maxDonated
}

func (r *RoundInfo) DonatedPlayer(donated int64) []string {
	res := make([]string, 0)
	for _, donatedInfo := range r.DonatedInfoList {
		if donatedInfo.DonatedMoney == donated {
			res = append(res, donatedInfo.PlayerId)
		}
	}
	return res
}

func (r *RoundInfo) PunishmentPlayer(playerIdList []string) {
	playerIdMap := make(map[string]struct{}, len(playerIdList))
	for _, playerId := range playerIdList {
		playerIdMap[playerId] = struct{}{}
	}
	for _, donatedInfo := range r.DonatedInfoList {
		if _, ok := playerIdMap[donatedInfo.PlayerId]; ok {
			donatedInfo.PunishmentMoney = r.PublicOpinion
		}
	}
}

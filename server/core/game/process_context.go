package game

type ProcessContext struct {
	Round            int64
	PlayerMap        map[string]*Player // 该map初始化后便不会变动，所以无需加锁便可操作。
	CurrentRoundInfo *RoundInfo
	RoundHistory     []*RoundInfo
	EndGame          chan struct{}
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
		hero := player.Hero
		var currentMoney, moneyLimit int64
		var heroName string
		if hero != nil {
			currentMoney, moneyLimit, heroName = hero.GetCurrentMoney(), hero.GetMoneyLimit(), hero.GetName()
		}
		p.CurrentRoundInfo.DonatedInfoList = append(p.CurrentRoundInfo.DonatedInfoList, &DonatedInfo{
			PlayerId:        playerId,
			Username:        player.Username,
			CurrentMoney:    currentMoney,
			HeroName:        heroName,
			MoneyLimit:      moneyLimit,
			DonatedMoney:    p.CurrentRoundInfo.defaultDonatedMoney(), // 当玩家超时或者掉线的时候，捐赠额度为舆论惩罚+2，对局势影响会较小
			PunishmentMoney: 0,                                        // 显式声明，表示全部字段都处理过了
			Bankrupt:        false,                                    // 同上
		})
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

func (p *ProcessContext) roundToHistory() {
	p.RoundHistory = append(p.RoundHistory, p.CurrentRoundInfo)
}

func (p *ProcessContext) judgementGameEnd() {
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

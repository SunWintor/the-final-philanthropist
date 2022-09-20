package hero

import "github.com/SunWintor/tfp/server/model"

type Hero interface {
	Init()

	GetHeroInfo() heroInfo
	GetMoneyLimit() int64
	GetCurrentMoney() int64
	GetName() string

	ToReply() *model.Hero

	IsBankrupt() bool
	DecMoney(int64)
	AddMoney(int64)

	OnRoundStart(ctx *SkillContext)
	OnRoundEnd(ctx *SkillContext)
	OnDonated(ctx *SkillContext)
	OnDonatedEnd(ctx *SkillContext)
}

type heroInfo struct {
	Id           int64
	Rarity       Rarity
	MoneyLimit   int64
	CurrentMoney int64
	Name         string
	Motto        string
	SkillInfo    []SkillInfo
}

type SkillInfo struct {
	Name string
	Desc string
}

type Rarity string

const (
	NRarity   Rarity = "N"
	RRarity   Rarity = "R"
	SRRarity  Rarity = "SR"
	SSRRarity Rarity = "SSR"
)

func (f *heroInfo) IsBankrupt() bool {
	return f.CurrentMoney <= 0
}

func (f *heroInfo) DecMoney(money int64) {
	f.AddMoney(-money)
}

func (f *heroInfo) AddMoney(money int64) {
	f.CurrentMoney += money
}

func (f *heroInfo) GetHeroInfo() heroInfo {
	return *f
}

func (f *heroInfo) GetMoneyLimit() int64 {
	return f.MoneyLimit
}

func (f *heroInfo) GetCurrentMoney() int64 {
	return f.CurrentMoney
}

func (f *heroInfo) GetName() string {
	return f.Name
}

func (f *heroInfo) ToReply() *model.Hero {
	return &model.Hero{
		MoneyLimit:   f.MoneyLimit,
		CurrentMoney: f.CurrentMoney,
		Name:         f.Name,
	}
}

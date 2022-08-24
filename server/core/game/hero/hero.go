package hero

import "github.com/SunWintor/tfp/server/model"

type Hero interface {
	Init()

	GetMoneyLimit() int64
	GetCurrentMoney() int64
	GetName() string
	GetSkillInfo() string

	ToReply() *model.Hero

	IsBankrupt() bool
	DecMoney(int64)
}

type heroInfo struct {
	MoneyLimit   int64
	CurrentMoney int64
	Name         string
	SkillInfo    string
}

func (f *heroInfo) IsBankrupt() bool {
	return f.CurrentMoney <= 0
}

func (f *heroInfo) DecMoney(money int64) {
	f.CurrentMoney -= money
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

func (f *heroInfo) GetSkillInfo() string {
	return f.SkillInfo
}

func (f *heroInfo) ToReply() *model.Hero {
	return &model.Hero{
		MoneyLimit:   f.MoneyLimit,
		CurrentMoney: f.CurrentMoney,
		Name:         f.Name,
		SkillInfo:    f.SkillInfo,
	}
}

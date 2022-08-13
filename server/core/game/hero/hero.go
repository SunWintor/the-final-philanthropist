package hero

import "github.com/SunWintor/tfp/server/model"

type Hero interface {
	Init()
	ToReply() *model.Hero
}

type heroInfo struct {
	MoneyLimit   int64
	CurrentMoney int64
	Name         string
	SkillInfo    string
}

func (f *heroInfo) ToReply() *model.Hero {
	return &model.Hero{
		MoneyLimit:   f.MoneyLimit,
		CurrentMoney: f.CurrentMoney,
		Name:         f.Name,
		SkillInfo:    f.SkillInfo,
	}
}

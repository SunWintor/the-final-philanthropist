package hero

type Foo2Die struct {
	heroInfo
}

func (f *Foo2Die) Init() {
	f.MoneyLimit = 108
	f.CurrentMoney = f.MoneyLimit
	f.Name = "富二代"
	f.SkillInfo = "胜利后变成富一代。"
}

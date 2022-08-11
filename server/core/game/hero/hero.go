package hero

type Hero interface {
	Init()
}

type heroInfo struct {
	MoneyLimit   int64
	CurrentMoney int64
	Name         string
	SkillInfo    string
}

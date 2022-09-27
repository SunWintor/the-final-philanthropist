package hero

type SkillFunc func(ctx *SkillContext)

type SkillContext struct {
	BankruptcyCount int64
}

func (f heroInfo) OnRoundStart(ctx *SkillContext) {
	return
}
func (f heroInfo) OnRoundEnd(ctx *SkillContext) {
	return
}
func (f heroInfo) OnDonated(donatedMoney int64) int64 {
	return donatedMoney
}
func (f heroInfo) OnDonatedEnd(ctx *SkillContext) {
	return
}

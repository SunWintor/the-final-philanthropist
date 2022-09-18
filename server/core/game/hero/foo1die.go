package hero

type Foo1Die struct {
	heroInfo
}

func (f *Foo1Die) Init() {
	f.Id = 15
	f.Rarity = SSRRarity
	f.MoneyLimit = 1
	f.CurrentMoney = f.MoneyLimit
	f.Name = "负一代"
	f.Motto = "可怜了我那还在襁褓中的孩子"
	f.SkillInfo = []SkillInfo{{
		Name: "负债",
		Desc: "被淘汰后使用富零代继续游戏"},
	}
}

package hero

type Full1Die struct {
	heroInfo
}

func (f *Full1Die) Init() {
	f.Id = 13
	f.Rarity = RRarity
	f.MoneyLimit = 110
	f.CurrentMoney = f.MoneyLimit
	f.Name = "富一代"
	f.Motto = "可惜父亲见不到我现在的样子了"
	f.SkillInfo = []SkillInfo{{
		Name: "白手起家",
		Desc: "使用富一代取得游戏胜利后\n该角色升级为富零代"},
	}
}

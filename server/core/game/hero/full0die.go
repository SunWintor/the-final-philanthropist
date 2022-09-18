package hero

type Full0Die struct {
	heroInfo
}

func (f *Full0Die) Init() {
	f.Id = 14
	f.Rarity = SRRarity
	f.MoneyLimit = 5
	f.CurrentMoney = f.MoneyLimit
	f.Name = "富零代"
	f.Motto = "孩子读书的钱一定要攒够"
	f.SkillInfo = []SkillInfo{{
		Name: "传宗",
		Desc: "被淘汰后使用富一代继续游戏"}, {
		Name: "接代",
		Desc: "使用富零代取得游戏胜利后\n该角色升级为负一代"},
	}
}

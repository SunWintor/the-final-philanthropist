package ecode

var (
	RoundAlreadyEndError = newErr(20000, "回合已结束", "round already end")
	StageNotDonatedError = newErr(20001, "未处于捐赠阶段", "stage not donated")
	PlayerNotExistsError = newErr(20002, "玩家不存在", "player not exists")
	GameNotStartError    = newErr(20003, "游戏未开始", "game not start")
	DonatedNotValidError = newErr(20004, "捐赠额度不能为负数", "donated not valid")
)

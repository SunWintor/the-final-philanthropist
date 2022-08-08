package ecode

var (
	AccountAlreadyRegError = newErr(-1000, "该账号已被注册", "account exists")
	AccountNotExistsError  = newErr(-1001, "账号不存在", "account not exists")
	PasswordError          = newErr(-1002, "密码错误", "password error")
)

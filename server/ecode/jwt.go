package ecode

var (
	TokenExpired     = newErr(10000, "登录已过期", "Token is expired")
	TokenNotValidYet = newErr(10001, "请稍后再试", "Token not active yet")
	TokenMalformed   = newErr(10002, "请登录", "That's not even a token")
	TokenInvalid     = newErr(10003, "登录异常", "Couldn't handle this token:")
)

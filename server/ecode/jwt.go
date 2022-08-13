package ecode

var (
	TokenExpired     = newErr(2000, "登录已过期", "Token is expired")
	TokenNotValIdYet = newErr(2001, "请稍后再试", "Token not active yet")
	TokenMalformed   = newErr(2002, "请登录", "That's not even a token")
	TokenInvalId     = newErr(2003, "登录异常", "Couldn't handle this token:")
)

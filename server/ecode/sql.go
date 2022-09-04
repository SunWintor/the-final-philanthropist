package ecode

import (
	"strings"
)

var (
	SqlNoResultErr = newErr(3000, "sql查不到数据", "no rows in result set")
)

func IsSqlNoResultErr(err error) bool {
	return strings.Contains(err.Error(), SqlNoResultErr.error.Error())
}

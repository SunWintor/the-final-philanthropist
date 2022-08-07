package gin_util

import (
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
)

func Result(c *gin.Context, code int64, data interface{}, msg string) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func SuccessWithMsg(c *gin.Context, msg string) {
	Result(c, SUCCESS, msg, "")
}

func FailWithError(c *gin.Context, err error) {
	if ecode, ok := ecode.ToEcode(err); ok {
		FailWithEcode(c, ecode)
		return
	}
	Result(c, ecode.UnknownError.Code, map[string]interface{}{}, err.Error())
}

func FailWithEcode(c *gin.Context, e *ecode.Ecode) {
	Result(c, e.Code, map[string]interface{}{}, e.Error())
}

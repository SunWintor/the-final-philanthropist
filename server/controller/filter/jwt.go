package filter

import (
	"github.com/SunWintor/tfp/server/common"
	http2 "github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			http2.FailWithEcode(c, ecode.NotLoginErr)
			c.Abort()
			return
		}
		j := common.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			http2.FailWithError(c, err)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

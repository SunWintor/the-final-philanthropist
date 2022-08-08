package filter

import (
	"github.com/SunWintor/tfp/server/common"
	http2 "github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/gin-gonic/gin"
)

func JWTAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		claims, err := JWTAuth(token)
		if err != nil {
			http2.FailWithError(c, err)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func JWTAuth(token string) (*common.TFPClaims, error) {
	if token == "" {
		return nil, ecode.NotLoginErr
	}
	j := common.NewJWT()
	return j.ParseToken(token)
}

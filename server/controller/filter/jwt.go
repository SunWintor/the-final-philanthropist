package filter

import (
	"github.com/SunWintor/tfp/server/common"
	http2 "github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/core/user"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/gin-gonic/gin"
	"strconv"
)

func JWTAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("x-Token")
		if err != nil {
			http2.FailWithError(c, err)
			c.Abort()
			return
		}
		claims, err := JWTAuth(token.Value)
		if err != nil {
			http2.FailWithError(c, err)
			c.Abort()
			return
		}
		err = userAuth(c, token.Value)
		if err != nil {
			http2.FailWithError(c, err)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func userAuth(c *gin.Context, token string) (err error) {
	userIdStr, err := c.Request.Cookie("user_id")
	if err != nil {
		return
	}
	userId, err := strconv.Atoi(userIdStr.Value)
	if err != nil {
		return
	}
	userInfo, ok := user.GetCopyById(int64(userId))
	if !ok {
		err = ecode.NotLoginErr
		return
	}
	if userInfo.Token != token {
		err = ecode.PermissionDeniedErr
		return
	}
	return
}

func JWTAuth(token string) (*common.TFPClaims, error) {
	if token == "" {
		return nil, ecode.NotLoginErr
	}
	j := common.NewJWT()
	return j.ParseToken(token)
}

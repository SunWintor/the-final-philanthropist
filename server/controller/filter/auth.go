package filter

import (
	"github.com/gin-gonic/gin"
)

// fixme 防止用户评级越权，先放在这，后面再完善
func EqualUltraViresHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.Request.Header.Get("x-token")
		//uId := c.Request.Header.Get("uid")
		// todo 如果用户不存在，直接返回让其重新登录。
		c.Next()
	}
}

package filter

import (
	"github.com/gin-gonic/gin"
)

// fixme 防止用户评级越权，先放在这，后面再完善
func EqualUltraViresHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.Request.Header.Get("x-token")
		//uid := c.Request.Header.Get("uid")
		c.Next()
	}
}

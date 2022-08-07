package controller

import "github.com/gin-gonic/gin"

func handleRoom(r *gin.Engine) {
	room := r.Group("/room")

	room.GET("/info", homeInfo)
}

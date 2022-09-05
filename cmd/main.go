package main

import (
	"flag"
	"fmt"
	"github.com/SunWintor/tfp/server/core/room"
	"time"

	"github.com/SunWintor/tfp/configs"
	"github.com/SunWintor/tfp/server/controller"
	"github.com/SunWintor/tfp/server/service"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
)

func setupRouter() *gin.Engine {
	s := service.New()
	room.Init(s)
	r := gin.New()
	r.Use(ginglog.Logger(3 * time.Second))
	r.Use(gin.Recovery())
	return controller.Handle(r, s)
}

func main() {
	flag.Parse()
	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", configs.GetConf().TFPServerPort))
}

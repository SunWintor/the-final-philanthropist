package main

import (
	"flag"
	"fmt"
	"github.com/SunWintor/tfp/configs"
	"github.com/SunWintor/tfp/server/controller"
	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"time"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(ginglog.Logger(3 * time.Second))
	r.Use(gin.Recovery())
	return controller.Handle(r)
}

func main() {
	flag.Parse()
	r := setupRouter()
	r.Run(fmt.Sprintf(":%s", configs.GetConf().TFPServerPort))
}

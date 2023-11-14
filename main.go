package main

import (
	"fmt"
	"github.com/device-server/global"
	"github.com/device-server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler.Register(r)
	err := r.Run(fmt.Sprintf(":%d", global.Cfg.ServerCfg.Port))
	if err != nil {
		panic(err.Error())
	}
}

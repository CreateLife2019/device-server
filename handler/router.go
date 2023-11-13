package handler

import (
	"github.com/device-server/handler/log"
	"github.com/device-server/handler/login"
	"github.com/device-server/handler/user"
	"github.com/device-server/handler/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func corsSetup(router *gin.Engine) bool {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Referer", "User-Agent", "Accept", "Authorization"}
	config.AllowAllOrigins = true
	corsHandler := cors.New(config)
	router.Use(corsHandler)

	return true
}

func Register(e *gin.Engine) {
	if gin.Mode() == "debug" {
		corsSetup(e)
	}
	web.Register(e)
	user.Register(e)
	log.Register(e)
	login.Register(e)
}

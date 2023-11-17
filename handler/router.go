package handler

import (
	"github.com/device-server/controller"
	"github.com/device-server/handler/account"
	"github.com/device-server/handler/group"
	"github.com/device-server/handler/log"
	"github.com/device-server/handler/login"
	"github.com/device-server/handler/middleware"
	"github.com/device-server/handler/proxy"
	"github.com/device-server/handler/tcp_client"
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
	mid := middleware.Register()
	g := e.Group("/device")
	g.POST("/logout", mid.LogoutHandler)
	g.POST("/refresh-token", mid.RefreshHandler)
	g.Use(mid.MiddlewareFunc())
	login.Register(e)
	web.Register(e)
	user.Register(g)
	group.Register(g)
	proxy.Register(g)
	log.Register(g)
	account.Register(g)
	controller.GetInstance().StartTcpServer(tcp_client.Onmessage, tcp_client.OnConnectionClose)
}

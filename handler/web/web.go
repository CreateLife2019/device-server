package web

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	e.NoRoute(static.Serve("/", static.LocalFile("./webui", false)))
}

package main

import (
	"github.com/device-server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler.Register(r)
	err := r.Run(":5999")
	if err != nil {
		panic(err.Error())
	}
}

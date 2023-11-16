package log

import (
	"github.com/device-server/domain/base"
	http2 "github.com/device-server/domain/request/http"
	http3 "github.com/device-server/domain/response/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.RouterGroup) {
	e.GET("/log/login", loginLogList)
}

// 登陆日志
func loginLogList(c *gin.Context) {
	loginReq := http2.LoginLogRequest{}
	if err := c.ShouldBindQuery(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.LoginLogResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {

		c.JSON(http.StatusOK, http3.LoginLogResponse{})
	}
}

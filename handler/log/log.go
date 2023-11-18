package log

import (
	"github.com/device-server/controller"
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
	listReq := http2.LoginLogRequest{}
	if err := c.ShouldBindQuery(&listReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.LoginLogResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})

	} else {

		var resp http3.LoginLogResponse
		resp, err = controller.GetInstance().AccountService().SearchLoginLog(listReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.LoginLogResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}

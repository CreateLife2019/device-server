package user

import (
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.Engine) {
	e.GET("/user", userList)
}

func userList(c *gin.Context) {
	loginReq := request.LoginLogRequest{}
	if err := c.ShouldBindQuery(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, response.LoginLogResponse{BaseResponse: response.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {

		c.JSON(http.StatusOK, response.LoginLogResponse{})
	}
}

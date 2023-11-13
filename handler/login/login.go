package login

import (
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.Engine) {
	e.POST("/login", login)
	e.POST("/verify-code", verifyCode)
}

// 登陆
func login(c *gin.Context) {
	loginReq := request.Login{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, response.LoginResponse{response.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {

		c.JSON(http.StatusOK, response.LoginResponse{})
	}
}

// 验证码
func verifyCode(c *gin.Context) {
	// 生成一个验证码
}

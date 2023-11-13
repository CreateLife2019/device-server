package login

import (
	"github.com/device-server/controller"
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
	loginReq := request.LoginRequest{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, response.LoginResponse{response.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = response.LoginResponse{}
		resp, err = controller.GetInstance().LoginService().Login(loginReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.LoginResponse{response.BaseResponse{
				Code: "500",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}

	}
}

// 验证码
func verifyCode(c *gin.Context) {
	// 生成一个验证码
}

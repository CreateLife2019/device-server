package login

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/device-server/controller"
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func Register(e *gin.Engine) {
	e.POST("/login", login)
	e.POST("/login/verify-code", verifyCode)
}

// 登陆
func login(c *gin.Context) {
	loginReq := request.LoginRequest{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, response.LoginResponse{BaseResponse: response.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = response.LoginResponse{}
		resp, err = controller.GetInstance().LoginService().Login(loginReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.LoginResponse{BaseResponse: response.BaseResponse{
				Code: "500",
				Msg:  err.Error(),
			}})
		} else {
			err = controller.GetInstance().VerifyCodeService().Check(loginReq.RequestId, loginReq.VerifyCode)
			if err != nil {
				c.JSON(http.StatusInternalServerError, response.LoginResponse{BaseResponse: response.BaseResponse{
					Code: "500",
					Msg:  err.Error(),
				}})
			} else {
				c.JSON(http.StatusOK, resp)
			}
		}
	}
}

// 验证码
func verifyCode(c *gin.Context) {
	var resp = response.VerifyCodeResponse{}
	// 生成一个验证码
	node, err := snowflake.NewNode(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.LoginResponse{BaseResponse: response.BaseResponse{
			Code: "500",
			Msg:  err.Error(),
		}})
		return
	} else {
		resp.Data.RequestId = node.Generate().Int64()
		resp.Data.VerifyCode = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
		// 保存数据哭
		err = controller.GetInstance().VerifyCodeService().Save(resp.Data.RequestId, resp.Data.VerifyCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.LoginResponse{BaseResponse: response.BaseResponse{
				Code: "500",
				Msg:  err.Error(),
			}})
		} else {
			resp.Code = "200"
			resp.Msg = constants.MessageSuc
			c.JSON(http.StatusOK, resp)
		}
	}
}

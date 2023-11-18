package login

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/device-server/controller"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
	http2 "github.com/device-server/domain/request/http"
	http3 "github.com/device-server/domain/response/http"
	"github.com/device-server/global"
	"github.com/device-server/internal/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func Register(e *gin.Engine) {
	e.POST("/device/login", login)
	e.POST("/device/login/verify-code", verifyCode)
}

// 登陆
func login(c *gin.Context) {
	loginReq := http2.LoginRequest{}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.LoginResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.LoginResponse{}
		err = controller.GetInstance().VerifyCodeService().Check(loginReq.RequestId, loginReq.VerifyCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.LoginResponse{BaseResponse: base.BaseResponse{
				Code: "500",
				Msg:  err.Error(),
			}})
		} else {
			loginReq.ClientIp = c.ClientIP()
			resp, err = controller.GetInstance().AccountService().Login(loginReq)
			if err != nil {
				c.JSON(http.StatusInternalServerError, http3.LoginResponse{BaseResponse: base.BaseResponse{
					Code: "500",
					Msg:  err.Error(),
				}})
				return
			}
			resp.Data.Token = utils.MakeTokenString([]byte(global.Cfg.ServerCfg.Key), global.Cfg.ServerCfg.Algorithm, loginReq.Account, loginReq.Password)
			c.JSON(http.StatusOK, resp)
		}
	}
}

// 验证码
func verifyCode(c *gin.Context) {
	var resp = http3.VerifyCodeResponse{}
	// 生成一个验证码
	node, err := snowflake.NewNode(1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.LoginResponse{BaseResponse: base.BaseResponse{
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
			c.JSON(http.StatusInternalServerError, http3.LoginResponse{BaseResponse: base.BaseResponse{
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

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, http3.LogoutResponse{})
}

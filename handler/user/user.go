package user

import (
	"github.com/device-server/controller"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
	http2 "github.com/device-server/domain/request/http"
	"github.com/device-server/domain/request/tcp"
	http3 "github.com/device-server/domain/response/http"
	"github.com/device-server/handler/tcp_client"
	"github.com/device-server/internal/repository/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(e *gin.RouterGroup) {
	e.GET("/user", userList)
	e.POST("/user/config/proxy", setProxy)
	e.POST("/user/config/send-proxy/:userId", sendProxy)
	e.GET("/user/config", userConfigList)
}

func userList(c *gin.Context) {
	userListReq := http2.UserListRequest{}
	if err := c.ShouldBindQuery(&userListReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UserListResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})

	} else {
		var resp http3.UserListResponse
		resp, err = controller.GetInstance().UserService().List(userListReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.UserListResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}

func setProxy(c *gin.Context) {
	proxyRequest := http2.ProxyRequest{}
	if err := c.ShouldBindJSON(&proxyRequest); err != nil {
		c.JSON(http.StatusBadRequest, http3.SetProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp http3.SetProxyResponse
		resp, err = controller.GetInstance().UserService().SetProxy(proxyRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.SetProxyResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			var user *entity.User
			user, err = controller.GetInstance().UserService().Get(proxyRequest.UserId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, http3.SetProxyResponse{BaseResponse: base.BaseResponse{
					Code: "400",
					Msg:  err.Error(),
				}})
				return
			}
			if proxyRequest.Immediately {
				proxyReq := tcp.ProxyRequest{
					RequestType: constants.TcpSetProxy,
				}
				proxyReq.HttpToTcp(proxyRequest)
				tcp_client.SendMessage(user.Phone, &proxyReq)
			}

			c.JSON(http.StatusOK, resp)
		}
	}
}
func sendProxy(c *gin.Context) {
	id := c.Param("userId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.SendProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	var user *entity.User
	user, err = controller.GetInstance().UserService().Get(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.SendProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	var userConfig *entity.UserConfig
	userConfig, err = controller.GetInstance().UserService().GetUserConfig(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.SendProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	proxyReq := userConfig.ToTcpProxy()
	tcp_client.SendMessage(user.Phone, &proxyReq)
	c.JSON(http.StatusOK, http3.SendProxyResponse{BaseResponse: base.BaseResponse{
		Code: constants.Status200,
		Msg:  constants.MessageSuc,
	}})
}
func userConfigList(c *gin.Context) {
	userListReq := http2.UserConfigListRequest{}
	if err := c.ShouldBindQuery(&userListReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UserConfigInfoListResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp http3.UserConfigInfoListResponse
		resp, err = controller.GetInstance().UserService().ListUserConfig(userListReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.UserConfigInfoListResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}

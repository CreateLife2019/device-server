package user

import (
	"github.com/device-server/controller"
	"github.com/device-server/domain/base"
	http2 "github.com/device-server/domain/request/http"
	http3 "github.com/device-server/domain/response/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.Engine) {
	e.GET("/user", userList)
	e.GET("/user/config/agent", addAgent)
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

func addAgent(c *gin.Context) {

}

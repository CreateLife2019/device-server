package user

import (
	"github.com/device-server/controller"
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(e *gin.Engine) {
	e.GET("/user", userList)
	e.GET("/user/config/agent", addAgent)
}

func userList(c *gin.Context) {
	userListReq := request.UserListRequest{}
	if err := c.ShouldBindQuery(&userListReq); err != nil {
		c.JSON(http.StatusBadRequest, response.UserListResponse{BaseResponse: response.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp response.UserListResponse
		resp, err = controller.GetInstance().UserService().List(userListReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.UserListResponse{BaseResponse: response.BaseResponse{
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

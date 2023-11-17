package group

import (
	"github.com/device-server/controller"
	"github.com/device-server/domain/base"
	http2 "github.com/device-server/domain/request/http"
	http3 "github.com/device-server/domain/response/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(e *gin.RouterGroup) {
	e.POST("/group", createGroup)
	e.GET("/group", groupList)
	e.GET("/group/:groupId", groupDetail)
	e.PUT("/group/:groupId", updateGroup)
	e.DELETE("/group/:groupId", deleteGroup)
}

func createGroup(c *gin.Context) {
	createReq := http2.CreateGroupRequest{}
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.CreateGroupResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.CreateGroupResponse{}
		resp, err = controller.GetInstance().GroupService().CreateGroup(createReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func updateGroup(c *gin.Context) {
	id := c.Param("groupId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.UpdateGroupResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	updateReq := http2.UpdateGroupRequest{GroupId: idLong}
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UpdateGroupResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.UpdateGroupResponse{}
		resp, err = controller.GetInstance().GroupService().UpdateGroup(updateReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func groupList(c *gin.Context) {
	listReq := http2.GroupListRequest{}
	if err := c.ShouldBindQuery(&listReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.GroupInfoListResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp http3.GroupInfoListResponse
		resp, err = controller.GetInstance().GroupService().GroupList(listReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.GroupInfoListResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func deleteGroup(c *gin.Context) {
	id := c.Param("groupId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteGroupResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	resp, err := controller.GetInstance().GroupService().DeleteGroup(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteGroupResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
func groupDetail(c *gin.Context) {
	id := c.Param("groupId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.GroupDetailResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	resp, err := controller.GetInstance().GroupService().GroupDetail(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.GroupDetailResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

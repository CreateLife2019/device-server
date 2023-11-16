package proxy

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
	e.POST("/proxy", createProxy)
	e.GET("/proxy", proxyList)
	e.GET("/proxy/:proxyId", proxyDetail)
	e.PUT("/proxy/:proxyId", updateProxy)
	e.DELETE("/proxy/:proxyId", deleteProxy)
}

func createProxy(c *gin.Context) {
	createReq := http2.CreateProxyRequest{}
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.LoginResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.CreateProxyResponse{}
		resp, err = controller.GetInstance().ProxyService().CreateProxy(createReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func updateProxy(c *gin.Context) {
	id := c.Param("proxyId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	updateReq := http2.UpdateProxyRequest{ProxyId: idLong}
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UpdateProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.UpdateProxyResponse{}
		resp, err = controller.GetInstance().ProxyService().UpdateProxy(updateReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func proxyList(c *gin.Context) {
	listReq := http2.ProxyListRequest{}
	if err := c.ShouldBindQuery(&listReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UserListResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp http3.ProxyInfoListResponse
		resp, err = controller.GetInstance().ProxyService().ProxyList(listReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.ProxyInfoListResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func deleteProxy(c *gin.Context) {
	id := c.Param("proxyId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteProxyResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	resp, err := controller.GetInstance().ProxyService().DeleteProxy(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteAccountResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
func proxyDetail(c *gin.Context) {
	id := c.Param("proxyId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.ProxyDetailResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	resp, err := controller.GetInstance().ProxyService().ProxyDetail(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.ProxyDetailResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

package account

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
	e.POST("/account", createAccount)
	e.GET("/account", accountList)
	e.PUT("/account", updateAccount)
	e.DELETE("/account/:accountId", deleteAccount)
}

func createAccount(c *gin.Context) {
	createReq := http2.CreateAccountRequest{}
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.LoginResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.CreateAccountResponse{}
		resp, err = controller.GetInstance().AccountService().CreateAccount(createReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func updateAccount(c *gin.Context) {
	updateReq := http2.UpdateAccountRequest{}
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UpdateAccountResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp = http3.UpdateAccountResponse{}
		resp, err = controller.GetInstance().AccountService().UpdateAccount(updateReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func accountList(c *gin.Context) {
	accountListReq := http2.AccountListRequest{}
	if err := c.ShouldBindQuery(&accountListReq); err != nil {
		c.JSON(http.StatusBadRequest, http3.UserListResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		var resp http3.AccountListResponse
		resp, err = controller.GetInstance().AccountService().AccountList(accountListReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, http3.AccountListResponse{BaseResponse: base.BaseResponse{
				Code: "400",
				Msg:  err.Error(),
			}})
		} else {
			c.JSON(http.StatusOK, resp)
		}
	}
}
func deleteAccount(c *gin.Context) {
	id := c.Param("accountId")
	idLong, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteAccountResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
		return
	}
	resp, err := controller.GetInstance().AccountService().DeleteAccount(idLong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http3.DeleteAccountResponse{BaseResponse: base.BaseResponse{
			Code: "400",
			Msg:  err.Error(),
		}})
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

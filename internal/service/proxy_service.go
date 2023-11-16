package service

import (
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
)

type ProxyService interface {
	CreateProxy(request http.CreateProxyRequest) (resp http2.CreateProxyResponse, err error)
	UpdateProxy(request http.UpdateProxyRequest) (resp http2.UpdateProxyResponse, err error)
	ProxyList(request http.ProxyListRequest) (resp http2.ProxyInfoListResponse, err error)
	DeleteProxy(proxy int64) (resp http2.DeleteProxyResponse, err error)
	ProxyDetail(proxy int64) (resp http2.ProxyDetailResponse, err error)
}

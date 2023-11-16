package tcp

import (
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request/http"
)

type ProxyInfo struct {
	ProxyHost   string `json:"proxyHost"`
	ProxyPort   int    `json:"proxyPort"`
	ProxySecret string `json:"proxySecret"`
}
type ProxyRequest struct {
	RequestType string `json:"requestType"`
	ProxyInfo   []ProxyInfo
}

func (p *ProxyRequest) HttpToTcp(request http.ProxyRequest) {
	//for _, v := range request.Proxies {
	//	p.ProxyInfo = append(p.ProxyInfo, ProxyInfo{
	//		ProxyHost:   v.ProxyHost,
	//		ProxyPort:   v.ProxyPort,
	//		ProxySecret: v.ProxySecret,
	//	})
	//}
}

func (p *ProxyRequest) ProtocolType() string {
	return constants.TcpSetProxy
}

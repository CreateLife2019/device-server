package tcp

import "github.com/device-server/domain/constants"

type LoginRequest struct {
	LoginTime   string `json:"LoginTime"`
	AppVersion  string `json:"appVersion"`
	ClientIP    string `json:"clientIp"`
	NickName    string `json:"nickName"`
	Phone       string `json:"phone"`
	ProxyHost   string `json:"proxyHost"`
	ProxyType   string `json:"proxyType"`
	RequestType string `json:"requestType"`
	UserName    string `json:"userName"`
}

func (l *LoginRequest) ProtocolType() string {
	return constants.TcpLoginType
}

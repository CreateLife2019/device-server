package request

type ClientLoginRequest struct {
	AppVersion  string `json:"appVersion"`
	ClientIP    string `json:"clientIp"`
	NickName    string `json:"nickName"`
	Phone       string `json:"phone"`
	RequestType string `json:"requestType"`
	UserName    string `json:"userName"`
}

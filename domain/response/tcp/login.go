package tcp

import (
	"encoding/json"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
)

type LoginResponse struct {
	base.Head
}

func (l *LoginResponse) BuildSuc() []byte {
	l.Head.RequestType = constants.TcpLoginType
	l.Head.Code = constants.Status200
	l.Head.Msg = constants.MessageSuc
	data, _ := json.Marshal(l)
	return data
}

func (l *LoginResponse) BuildFailed(code string) []byte {
	l.Code = code
	l.Msg = "请求失败"
	data, _ := json.Marshal(l)
	return data
}

package tcp

import (
	"encoding/json"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
)

type HeartbeatResponse struct {
	base.Head
	UserId int64 `json:"userId"`
}

func (l *HeartbeatResponse) BuildSuc() []byte {
	l.Head.RequestType = constants.TcpLoginType
	l.Head.Code = constants.Status200
	l.Head.Msg = constants.MessageSuc
	data, _ := json.Marshal(l)
	return data
}

func (l *HeartbeatResponse) BuildFailed(code string) []byte {
	l.Code = code
	l.Msg = "请求失败"
	data, _ := json.Marshal(l)
	return data
}

package tcp

import (
	"encoding/json"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
)

type OfflineResponse struct {
	base.Head
	Phone string `json:"phone"`
}

func (l *OfflineResponse) BuildSuc() []byte {
	l.Head.RequestType = constants.TcpOffline
	l.Head.Code = constants.Status200
	l.Head.Msg = constants.MessageSuc
	data, _ := json.Marshal(l)
	return data
}

func (l *OfflineResponse) BuildFailed(code string) []byte {
	l.Code = code
	l.Msg = "请求失败"
	data, _ := json.Marshal(l)
	return data
}

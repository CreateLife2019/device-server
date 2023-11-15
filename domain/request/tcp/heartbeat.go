package tcp

import "github.com/device-server/domain/constants"

type HeartbeatRequest struct {
	RequestType string `json:"requestType"`
	Status      string `json:"status"`
	Phone       string `json:"phone"`
}

func (l *HeartbeatRequest) ProtocolType() string {
	return constants.TcpHeartbeat
}

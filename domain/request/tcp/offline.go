package tcp

import "github.com/device-server/domain/constants"

type OfflineRequest struct {
	Phone       string `json:"phone"`
	RequestType string `json:"requestType"`
	Status      string `json:"status"`
}

func (l *OfflineRequest) ProtocolType() string {
	return constants.TcpOffline
}

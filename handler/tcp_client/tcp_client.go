package tcp_client

import (
	"encoding/json"
	"github.com/device-server/controller"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
	"github.com/device-server/domain/request/tcp"
	"github.com/device-server/internal/tcp_server"
	"github.com/sirupsen/logrus"
)

func Onmessage(c *tcp_server.Client, message []byte) {
	head := base.Head{}
	err := json.Unmarshal(message, &head)
	if err != nil {
		logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
	} else {
		handleMessage(c, message, &head)
	}
}
func handleMessage(c *tcp_server.Client, message []byte, head *base.Head) {
	switch head.RequestType {
	case constants.TcpLoginType:
		logReq := tcp.LoginRequest{}
		err := json.Unmarshal(message, &logReq)
		if err != nil {
			logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
			return
		}
		resp := controller.GetInstance().UserService().Login(logReq)
		err = c.SendBytes(resp)
		if err != nil {
			logrus.Errorf("收到客户端消息，回复失败:%s", err.Error())
			return
		}
	case constants.TcpHeartbeat:
		logReq := tcp.HeartbeatRequest{}
		err := json.Unmarshal(message, &logReq)
		if err != nil {
			logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
			return
		}
		resp := controller.GetInstance().UserService().Heartbeat(logReq)
		err = c.SendBytes(resp)
		if err != nil {
			logrus.Errorf("收到客户端消息，回复失败:%s", err.Error())
			return
		}
	}

}

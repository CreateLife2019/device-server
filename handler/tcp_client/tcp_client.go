package tcp_client

import (
	"encoding/json"
	"github.com/device-server/controller"
	"github.com/device-server/domain/base"
	"github.com/device-server/domain/constants"
	tcpBase "github.com/device-server/domain/request/base"
	"github.com/device-server/domain/request/tcp"
	"github.com/device-server/internal/tcp_server"
	"github.com/sirupsen/logrus"
	"sync"
)

var clientMap = map[string]*tcp_server.Client{}
var lock sync.Locker = &sync.Mutex{}

func addClient(phone string, c *tcp_server.Client) {
	lock.Lock()
	defer lock.Unlock()
	clientMap[phone] = c
}
func getClient(phone string) *tcp_server.Client {
	lock.Lock()
	defer lock.Unlock()
	return clientMap[phone]
}
func deleteClient(c *tcp_server.Client) {
	lock.Lock()
	defer lock.Unlock()
	phone := ""
	for key, v := range clientMap {
		if v.Conn() == c.Conn() {
			phone = key
			break
		}
	}
	if phone != "" {
		delete(clientMap, phone)
	}
}

func Onmessage(c *tcp_server.Client, message []byte) {
	head := base.Head{}
	err := json.Unmarshal(message, &head)
	if err != nil {
		logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
	} else {
		handleMessage(c, message, &head)
	}
}
func OnConnectionClose(c *tcp_server.Client, err error) {
	deleteClient(c)
}
func handleMessage(c *tcp_server.Client, message []byte, head *base.Head) {
	var sendBytes = make([]byte, 0)
	switch head.RequestType {
	case constants.TcpLoginType:
		logReq := tcp.LoginRequest{}
		err := json.Unmarshal(message, &logReq)
		if err != nil {
			logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
			return
		}
		resp, err := controller.GetInstance().UserService().Login(logReq)
		if err != nil {
			sendBytes = resp.BuildFailed("500")
		} else {
			sendBytes = resp.BuildSuc()
		}
		err = c.SendBytes(sendBytes)
		if err != nil {
			logrus.Errorf("收到客户端消息，回复失败:%s", err.Error())
			return
		}
		addClient(logReq.Phone, c)
	case constants.TcpHeartbeat:
		logReq := tcp.HeartbeatRequest{}
		err := json.Unmarshal(message, &logReq)
		if err != nil {
			logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
			return
		}
		resp, err := controller.GetInstance().UserService().Heartbeat(logReq)
		if err != nil {
			sendBytes = resp.BuildFailed("500")
		} else {
			sendBytes = resp.BuildSuc()
		}
		err = c.SendBytes(sendBytes)
		if err != nil {
			logrus.Errorf("收到客户端消息，回复失败:%s", err.Error())
			return
		}
	case constants.TcpOffline:
		offlineReq := tcp.OfflineRequest{}
		err := json.Unmarshal(message, &offlineReq)
		if err != nil {
			logrus.Errorf("收到客户端消息，解析失败:%s", err.Error())
			return
		}
		resp, err := controller.GetInstance().UserService().Offline(offlineReq)
		if err != nil {
			sendBytes = resp.BuildFailed("500")
		} else {
			sendBytes = resp.BuildSuc()
		}
		err = c.SendBytes(sendBytes)
		if err != nil {
			logrus.Errorf("收到客户端消息，回复失败:%s", err.Error())
			return
		}
	}
}

func SendMessage(phone string, req tcpBase.TcpRequestProtocol) {
	client := getClient(phone)
	if client != nil {
		sendBytes, err := json.Marshal(req)
		if err != nil {
			logrus.Errorf("发送给客户端消息失败，:%s", err.Error())
			return
		}
		err = client.SendBytes(sendBytes)
		if err != nil {
			logrus.Errorf("发送给客户端消息失败，:%s", err.Error())
			return
		}
	}
}

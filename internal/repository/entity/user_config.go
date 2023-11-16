package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/device-server/domain/request/http"
	http2 "github.com/device-server/domain/response/http"
	"time"
)

type SessionInfo struct {
	// 开始时间
	StartTime time.Time `json:"startTime"`
	// 结束时间
	EndTime time.Time `json:"endTime"`
	// 类型 1 个人 2 群组
	SessionType int `json:"sessionType"`
	// 目标id
	DestId int64 `json:"destId"`
}
type Proxy struct {
	Host    string    `json:"host"`
	Port    int       `json:"port"`
	Timeout int       `json:"timeout"`
	Secret  string    `json:"secret"`
	Time    time.Time `json:"time"`
}
type Proxies []Proxy
type Sessions []SessionInfo
type UserConfig struct {
	Base
	MessageIntercept int      `json:"messageIntercept" gorm:"column:f_message_intercept"` // 消息拦截
	SyncSessions     Sessions `json:"syncMessage"gorm:"column:f_sync_sessions;json"`      // 会话同步
	UserId           int64    `json:"userId" gorm:"column:f_user_id"`                     // 用户id
	Proxies          Proxies  `json:"proxies" gorm:"column:f_proxies;json"`               // 代理
}

func (a *UserConfig) TableName() string {
	return "f_user_config"
}

const (
	empty = `{}`
)

func (p Proxies) Value() (driver.Value, error) {
	if len(p) == 0 {
		return empty, nil
	}
	return json.Marshal(p)
}

// Scan 实现方法
func (p *Proxies) Scan(data interface{}) error {
	b, ok := data.([]byte)
	if !ok {
		// Images - 占位
		b = []byte(empty)
	}
	return json.Unmarshal(b, p)
}
func (p Sessions) Value() (driver.Value, error) {
	if len(p) == 0 {
		return empty, nil
	}
	return json.Marshal(p)
}

// Scan 实现方法
func (p *Sessions) Scan(data interface{}) error {
	b, ok := data.([]byte)
	if !ok {
		// Images - 占位
		b = []byte(empty)
	}
	return json.Unmarshal(b, p)
}
func (u *UserConfig) ReadProxy(request http.ProxyRequest) {
	item := Proxy{}
	for _, v := range request.Proxies {
		item.Host = v.ProxyHost
		item.Secret = v.ProxySecret
		item.Port = v.ProxyPort
		item.Time = time.Now()
		u.Proxies = append(u.Proxies, item)
	}
	u.UserId = request.UserId
}
func (u *UserConfig) ToResponse() http2.UserConfigInfo {
	resp := http2.UserConfigInfo{
		Id:     u.Id,
		UserId: u.UserId,
	}
	for _, v := range u.Proxies {
		resp.Proxies = append(resp.Proxies, http2.ProxyInfo{
			ProxyHost:   v.Host,
			ProxyPort:   v.Port,
			ProxySecret: v.Secret,
			SetTime:     v.Time,
		})
	}

	return resp

}

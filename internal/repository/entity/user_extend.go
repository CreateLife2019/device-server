package entity

import "time"

type UserExtend struct {
	Base
	UserId        int64     `json:"userId" gorm:"index:idx_user_id,unique;column:f_user_id;NOT NULL" `
	Online        int       `json:"online" gorm:"column:f_online;NOT NULL;default:2" ` // 1在线， 2 离线，
	ClientIp      string    `json:"clientIp" gorm:"column:f_client_ip;NOT NULL;default:''"`
	LoginTime     time.Time `json:"loginTime" gorm:"column:f_login_time"`
	AppVersion    string    `json:"appVersion" gorm:"column:f_app_version"`
	ProxyType     string    `json:"proxyType" gorm:"column:f_proxy_type"`
	HeartbeatTime time.Time `json:"heartbeatTime" gorm:"column:f_heartbeat_time"`
}

func (a *UserExtend) TableName() string {
	return "t_user_extend"
}

package entity

import "time"

type Proxy struct {
	Base
	Host    string    `json:"host" gorm:"column:f_host;NOT NULL;default:''"`
	Port    int       `json:"port" gorm:"column:f_port;NOT NULL;default:0" `
	Timeout int       `json:"timeout" gorm:"column:f_timeout;NOT NULL;default:0"`
	Secret  string    `json:"secret" gorm:"column:f_secret;NOT NULL;default:''"`
	Time    time.Time `json:"time" gorm:"-"`
}

func (a *Proxy) TableName() string {
	return "t_proxy"
}

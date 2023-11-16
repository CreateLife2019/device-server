package entity

type LoginLog struct {
	Base
	AccountId int64  `json:"accountId" gorm:"index:idx_account_id;column:f_account_id;NOT NULL" `
	Account   string `gorm:"column:f_account;NOT NULL"` // 用户名
	Ip        string `gorm:"column:f_ip;NOT NULL"`      // 用户名
}

func (l *LoginLog) TableName() string {
	return "t_login_log"
}

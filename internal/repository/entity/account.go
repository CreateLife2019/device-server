package entity

import "time"

// 后端管理账号
type Account struct {
	Base
	UserName  string    `gorm:"column:f_user_name;NOT NULL"` // 用户名
	Password  string    `gorm:"column:f_password;NOT NULL"`  // 密码
	LoginTime time.Time `gorm:"column:f_login_time;"`
}

func (a *Account) TableName() string {
	return "t_account"
}

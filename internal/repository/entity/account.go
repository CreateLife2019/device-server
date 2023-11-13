package entity

// 后端管理账号
type Account struct {
	Base
	UserName string `gorm:"column:f_user_name;NOT NULL"` // 用户名
	Password string `gorm:"column:f_password;NOT NULL"`  // 密码
}

func (a *Account) TableName() string {
	return "t_account"
}

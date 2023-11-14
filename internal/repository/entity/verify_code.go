package entity

type VerifyCode struct {
	Base
	RequestId  int64  `gorm:"index:idx_request_id,unique;column:f_request_id;NOT NULL"` // 用户名
	VerifyCode string `gorm:"column:f_verify_code;NOT NULL"`                            // 密码
}

func (a *VerifyCode) TableName() string {
	return "t_verify_code"
}

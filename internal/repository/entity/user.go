package entity

type User struct {
	Base
	Name     string `json:"name" gorm:"column:f_name"`                                   // 名称
	Phone    string `json:"phone" gorm:"index:idx_phone,unique;column:f_phone;NOT NULL"` // 电话
	NickName string `json:"nickName" gorm:"column:f_nick_name"`                          // 用户名
	Avatar   string `json:"avatar" gorm:"column:f_avatar"`                               // 头像
	Remark   string `json:"remark" gorm:"column:f_remark"`                               // 备注信息
}

func (a *User) TableName() string {
	return "t_user"
}

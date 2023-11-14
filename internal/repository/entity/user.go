package entity

type User struct {
	Base
	Name     string `json:"name" gorm:"f_name"`          // 名称
	Phone    string `json:"phone" gorm:"f_phone"`        // 电话
	NickName string `json:"nickName" gorm:"f_nick_name"` // 用户名
	Avatar   string `json:"avatar" gorm:"f_avatar"`      // 头像
	Remark   string `json:"remark" gorm:"f_remark"`      // 备注信息
}

func (a *User) TableName() string {
	return "t_user"
}

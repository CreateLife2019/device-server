package entity

type User struct {
	Base
	Name     string `json:"name"`     // 名称
	Phone    string `json:"phone"`    // 电话
	NickName string `json:"nickName"` // 用户名
	Avatar   string `json:"avatar"`   // 头像
	Remark   string `json:"remark"`   // 备注信息
}

func (a *User) TableName() string {
	return "t_user"
}

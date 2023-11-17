package entity

type UserGroup struct {
	Base
	UserId  int64 `json:"userId" gorm:"column:f_user_id"`
	GroupId int64 `json:"groupId" gorm:"column:f_group_id"`
}

func (a *UserGroup) TableName() string {
	return "t_user_group"
}

package entity

type Group struct {
	Base
	Name   string `json:"name" gorm:"index:idx_group_name,unique;column:f_name;NOT NULL;default:''"`
	System int    `json:"system" gorm:"column:f_system;NOT NULL;default:2"` // 系统，默认都不是系统分组
}

func (a *Group) TableName() string {
	return "f_group"
}

package entity

type Group struct {
	Base
	Name string `json:"name" gorm:"index:idx_group_name,unique;column:f_name;NOT NULL;default:''"`
}

func (a *Group) TableName() string {
	return "f_group"
}

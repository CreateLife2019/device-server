package entity

import "time"

type Base struct {
	Id        int64      `gorm:"column:f_id;primaryKey" db:"f_id" json:"id" form:"f_id"`                                                          // 自增id
	CreatedAt *time.Time `gorm:"column:f_created_at;default:current_timestamp"  json:"createdAt" form:"f_created_at"`                             // 创建时间
	UpdatedAt *time.Time `gorm:"column:f_updated_at;default:current_timestamp on update current_timestamp"  json:"updatedAt" form:"f_updated_at"` // 更新时间
}

type Page struct {
	Page     int64
	PageSize int64
	Total    int64
}

func (p *Page) Paginate() (start, end int64) {
	start = (p.Page - 1) * p.PageSize
	end = start + p.PageSize
	if start > p.Total {
		start = p.Total
		end = p.Total
		return
	}
	if end > p.Total {
		end = p.Total
	}
	return
}

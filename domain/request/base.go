package request

type PageInfo struct {
	Page     int64 `form:"page,omitempty" binding:"required"`
	PageSize int64 `form:"pageSize,omitempty" binding:"required"`
}

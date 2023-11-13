package request

type PageInfo struct {
	PageNo   int64 `form:"PageNo,omitempty" binding:"required"`
	PageSize int64 `form:"PageSize,omitempty" binding:"required"`
}

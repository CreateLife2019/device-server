package base

type PageInfo struct {
	Page     int64 `form:"page,omitempty" binding:"required" json:"page"`
	PageSize int64 `form:"pageSize,omitempty" binding:"required" json:"pageSize"`
	Total    int64 `json:"total"`
}
type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

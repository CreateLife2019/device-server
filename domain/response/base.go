package response

type BaseResponse struct {
	Code string `json:"Code"`
	Msg  string `json:"Msg"`
}

type PageInfo struct {
	PageNo   int64 `form:"pageNo" binding:"required"`
	PageSize int64 `form:"pageSize" binding:"required"`
	Total    int64 `form:"total" binding:"required"`
}

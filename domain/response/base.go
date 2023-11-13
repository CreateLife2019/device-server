package response

type BaseResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type PageInfo struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	Total    int64 `json:"total"`
}

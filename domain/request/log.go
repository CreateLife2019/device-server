package request

type LoginLogRequest struct {
	PageInfo
	UserId int64 `form:"UserId"`
}

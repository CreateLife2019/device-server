package response

type AdminIdData struct {
	Id int64 `json:"id"`
}
type LoginResponse struct {
	BaseResponse
	Data AdminIdData `json:"data"`
}

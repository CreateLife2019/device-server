package response

type VerifyCodeData struct {
	RequestId  int64  `json:"requestId"`
	VerifyCode string `json:"verifyCode"`
}
type VerifyCodeResponse struct {
	BaseResponse
	Data VerifyCodeData `json:"data"`
}

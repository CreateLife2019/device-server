package response

type VerifyCodeResponse struct {
	Id   int64  `json:"Id"`
	Code string `json:"Code"`
}

package base

type Head struct {
	RequestType string `json:"requestType"`
	Code        string `json:"code"`
	Msg         string `json:"msg"`
}

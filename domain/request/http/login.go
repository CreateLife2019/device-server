package http

type LoginRequest struct {
	Account    string `form:"account" json:"account" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	VerifyCode string `form:"verifyCode" json:"verifyCode" binding:"required"`
	RequestId  int64  `form:"requestId" json:"requestId,string" binding:"required"`
}

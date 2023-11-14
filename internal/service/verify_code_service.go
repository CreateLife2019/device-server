package service

type VerifyCodeService interface {
	Check(requestId int64, code string) (err error)
	Save(requestId int64, code string) (err error)
}

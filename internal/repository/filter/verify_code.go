package filter

import "gorm.io/gorm"

func WithVerifyCode(requestId int64, code string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_request_id = ? and f_verify_code = ? ", requestId, code)
		return db
	}
}

package filter

import "gorm.io/gorm"

func WithPhone(phone string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_phone = ?  ", phone)
		return db
	}
}

func WithUserId(userId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_user_id = ?  ", userId)
		return db
	}
}
func WithOnline() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_online = 1  ")
		return db
	}
}
func WithId(userId int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_id = ?  ", userId)
		return db
	}
}
func WithInId(userId []int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_id in ?  ", userId)
		return db
	}
}

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
		if userId != 0 {
			db = db.Where("f_user_id = ?  ", userId)
		}

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
		if userId != 0 {
			db = db.Where("f_id = ?  ", userId)
		}

		return db
	}
}
func WithInId(ids []int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_id in ?  ", ids)
		return db
	}
}
func WithInUserId(userId []int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_user_id in ?  ", userId)
		return db
	}
}
func WithGroupId(ids []int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		db = db.Where("f_group_id in ?  ", ids)
		return db
	}
}
func WithLickName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {
			condition := "%" + name + "%"
			db = db.Where("f_name like ?  ", condition)
		}

		return db
	}
}
func WithName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != "" {

			db = db.Where("f_name = ?  ", name)
		}

		return db
	}
}

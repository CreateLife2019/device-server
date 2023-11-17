package persistence

import (
	"github.com/device-server/internal/repository/entity"
	"gorm.io/gorm"
)

type UserIer interface {
	SearchUser(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.User, err error)
	SaveOrUpdate(db *gorm.DB, user *entity.User) error
	SaveOrUpdateExtend(db *gorm.DB, user *entity.UserExtend) error
	Get(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.User, err error)
	UpdateUserExtend(tx *gorm.DB, in *entity.UserExtend, scopes ...func(db *gorm.DB) *gorm.DB) (err error)
	Update(tx *gorm.DB, in *entity.User, scopes ...func(db *gorm.DB) *gorm.DB) (err error)
	SearchUserExtend(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.UserExtend, err error)
	GetOrCreateUserConfig(db *gorm.DB, in *entity.UserConfig, scopes ...func(*gorm.DB) *gorm.DB) (out *entity.UserConfig, err error)
	SearchUserConfig(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.UserConfig, err error)
	GetUserConfig(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) (account *entity.UserConfig, err error)
	UpdateUserConfig(tx *gorm.DB, in *entity.UserConfig, scopes ...func(db *gorm.DB) *gorm.DB) (err error)
	OfflineUsers(db *gorm.DB, userIds []int64) error
	SearchUserGroup(db *gorm.DB, page *entity.Page, scopes ...func(*gorm.DB) *gorm.DB) (users []*entity.UserGroup, err error)
}

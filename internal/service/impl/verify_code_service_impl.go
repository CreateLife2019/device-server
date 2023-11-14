package impl

import (
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/filter"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"gorm.io/gorm"
)

type VerifyCodeServiceImpl struct {
	db   *gorm.DB
	vIer persistence.VerifyCodeIer
}

func NewUserVerifyCodeService(db *gorm.DB) *VerifyCodeServiceImpl {
	return &VerifyCodeServiceImpl{db: db, vIer: &impl.VerifyCodeImpl{}}
}

func (v *VerifyCodeServiceImpl) Check(requestId int64, code string) (err error) {
	_, err = v.vIer.Get(v.db, filter.WithVerifyCode(requestId, code))
	if err != nil {
		return
	}
	return
}
func (v *VerifyCodeServiceImpl) Save(requestId int64, code string) (err error) {
	return v.vIer.BatchSave(v.db, []*entity.VerifyCode{&entity.VerifyCode{
		RequestId:  requestId,
		VerifyCode: code,
	}})
}

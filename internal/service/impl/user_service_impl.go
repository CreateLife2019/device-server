package impl

import (
	"github.com/device-server/domain/request"
	"github.com/device-server/domain/response"
	"github.com/device-server/internal/repository/entity"
	"github.com/device-server/internal/repository/persistence"
	"github.com/device-server/internal/repository/persistence/impl"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	db   *gorm.DB
	user persistence.UserIer
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db, user: &impl.UserIerImpl{}}
}
func (u *UserServiceImpl) List(request request.UserListRequest) (resp response.UserListResponse, err error) {
	users := make([]*entity.User, 0)
	page := &entity.Page{
		Page:     request.Page,
		PageSize: request.PageSize,
	}
	users, err = u.user.SearchUser(u.db, page)
	if err != nil {
		return
	}
	resp.Total = page.Total
	resp.Page = page.Page
	resp.PageSize = page.PageSize
	resp.Users = make([]response.UserInfo, 0)
	for _, v := range users {
		resp.Users = append(resp.Users, response.UserInfo{
			Name:       v.Name,
			Phone:      v.Phone,
			NickName:   v.NickName,
			Avatar:     v.Avatar,
			Remark:     v.Remark,
			DeviceName: "",
			Ip:         "",
			Online:     0,
			Agent:      0,
		})
	}
	return
}

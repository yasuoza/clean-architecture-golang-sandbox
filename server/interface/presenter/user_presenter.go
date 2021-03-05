package presenter

import (
	"github.com/yasuoza/clean-architecture-go/server/domain/model"
	"github.com/yasuoza/clean-architecture-go/server/usecase/presenter"
)

type userPresenter struct {
}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}

	return us
}

func (up *userPresenter) CreateUser(u *model.User) *model.User {
	u.Name = "CreatedUser: " + u.Name
	return u
}

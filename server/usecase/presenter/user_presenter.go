package presenter

import "github.com/yasuoza/clean-architecture-go/server/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
	CreateUser(u *model.User) *model.User
}

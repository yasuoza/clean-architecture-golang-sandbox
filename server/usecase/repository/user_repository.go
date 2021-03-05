package repository

import "github.com/yasuoza/clean-architecture-go/server/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

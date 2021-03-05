package interactor

import (
	"errors"

	"github.com/yasuoza/clean-architecture-go/server/domain/model"
	"github.com/yasuoza/clean-architecture-go/server/usecase/presenter"
	"github.com/yasuoza/clean-architecture-go/server/usecase/repository"
)

type userInteractor struct {
	DBRepository   repository.DBRepository
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}

func NewUserInteractor(d repository.DBRepository, r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{d, r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) Create(u *model.User) (*model.User, error) {
	data, err := us.DBRepository.Transaction(func(_ interface{}) (interface{}, error) {
		u, err := us.UserRepository.Create(u)

		return u, err
	})

	if !errors.Is(err, nil) {
		return nil, err
	}

	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	return us.UserPresenter.CreateUser(user), nil
}

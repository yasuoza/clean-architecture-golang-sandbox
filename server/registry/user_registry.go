package registry

import (
	"github.com/yasuoza/clean-architecture-go/server/interface/controller"
	ip "github.com/yasuoza/clean-architecture-go/server/interface/presenter"
	ir "github.com/yasuoza/clean-architecture-go/server/interface/repository"
	"github.com/yasuoza/clean-architecture-go/server/usecase/interactor"
	up "github.com/yasuoza/clean-architecture-go/server/usecase/presenter"
	ur "github.com/yasuoza/clean-architecture-go/server/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(
		r.NewDBRepository(),
		r.NewUserRepository(),
		r.NewUserPresenter(),
	)
}

func (r *registry) NewDBRepository() ur.DBRepository {
	return ir.NewDBRepository(r.db)
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

package user

import (
	"Latihan_1/app/domain/repository"
	"Latihan_1/app/domain/usecase"
)

type usecaseUser struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserUsecase(userRepo repository.UserRepositoryInterface) usecase.UserUsecaseInterface {
	return &usecaseUser{
		userRepository: userRepo,
	}
}

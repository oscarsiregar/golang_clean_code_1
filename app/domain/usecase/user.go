package usecase

import "Latihan_1/app/domain/entity"

type UserUsecaseInterface interface {
	CreateUser(us entity.User) (*entity.User, error)
}

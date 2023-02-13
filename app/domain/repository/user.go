package repository

import "Latihan_1/app/domain/entity"

type UserRepositoryInterface interface {
	CreateUser(user *entity.User) (*entity.User, error)
}

package repository

import (
	"Latihan_1/app/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

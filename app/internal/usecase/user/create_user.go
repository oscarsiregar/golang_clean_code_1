package user

import (
	"Latihan_1/app/domain/entity"
)

func (u *usecaseUser) CreateUser(us entity.User) (*entity.User, error) {
	// Validasi DTO
	err := us.ValidateData()
	if err != nil {
		return nil, err
	}

	// Buat Mapper
	newUser := entity.NewUser(us)

	// Create New User
	user, err := u.userRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

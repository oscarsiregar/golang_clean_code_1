package repository

import (
	"Latihan_1/app/domain/entity"
	"Latihan_1/app/internal/repository/psql/user/mapper"
)

func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	// Konversi Entity Ke Model
	mappedUser := mapper.ConvertEntityToModelUser(user)

	err := r.db.Create(mappedUser).Error

	mappedUserResult := mapper.ConvertModelToEntityUser(mappedUser)

	return mappedUserResult, err
}

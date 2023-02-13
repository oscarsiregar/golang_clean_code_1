package mapper

import (
	"Latihan_1/app/domain/entity"
	"Latihan_1/app/internal/repository/psql/user/model"

	"github.com/google/uuid"
)

func ConvertEntityToModelUser(user *entity.User) *model.User {
	convertedUser := model.User{
		ID:         user.ID.String(),
		Name:       user.Name,
		Age:        user.Age,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}
	return &convertedUser
}

func ConvertModelToEntityUser(user *model.User) *entity.User {
	id, _ := uuid.Parse(user.ID)
	convertedUser := entity.User{
		ID:         id,
		Name:       user.Name,
		Age:        user.Age,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}
	return &convertedUser
}

package delivery

import (
	"Latihan_1/app/domain/delivery"
	"Latihan_1/app/domain/usecase"
)

type userHandler struct {
	userUsecase usecase.UserUsecaseInterface
}

func NewUserHandler(userUsecase usecase.UserUsecaseInterface) delivery.UserDeliveryInterface {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

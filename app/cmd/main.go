package main

import (
	"Latihan_1/app/internal/config"
	userDelivery "Latihan_1/app/internal/delivery/user"
	userRepository "Latihan_1/app/internal/repository/psql/user"
	"Latihan_1/app/internal/repository/psql/user/model"
	userUsecase "Latihan_1/app/internal/usecase/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Latihan 1")
	// Koneksi ke db
	db, err := config.ConnectPostgreSQL()
	if err != nil {
		log.Fatal("Gagal Koneksi DB :", err.Error())
	}

	fmt.Printf("Success Connect DB : %s\n", "latihan_1")

	// Migrasi Model
	db.AutoMigrate(&model.User{})

	// Panggil user repository
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userDelivery := userDelivery.NewUserHandler(userUsecase)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/user", userDelivery.CreateUser)
	router.Run(":3000")
}

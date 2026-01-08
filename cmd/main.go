package main

import (
	"modulo/controller"
	"modulo/db"
	"modulo/repository"
	"modulo/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// repository
	UserRepository := repository.NewUserRepository(dbConnection)

	// usecase
	UserUsecase := usecase.NewUserUsecase(UserRepository)

	// controller
	UserController := controller.NewUserController(UserUsecase)

	server.GET("/users", UserController.GetUsers)

	server.Run(":8000")
}

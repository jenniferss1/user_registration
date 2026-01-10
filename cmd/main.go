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

	// routes
	server.GET("/users", UserController.GetUsers)
	server.POST("/user", UserController.CreateUser)
	server.GET("/user/:userId", UserController.GetUserByID)

	server.Run(":8000")
}

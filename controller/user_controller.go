package controller

import (
	"modulo/model"
	"modulo/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) userController {
	return userController{
		userUsecase: usecase,
	}
}

func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.userUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, users)
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var user model.Users
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedUser, err := u.userUsecase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)
}

func (u *userController) GetUserByID(ctx *gin.Context) {

	id := ctx.Param("userId")
	if id == "" {
		response := model.Response{
			Message: "user id cannot be empty.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "user id have to be a number.",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := u.userUsecase.GetUserByID(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if user == nil {
		response := model.Response{
			Message: "user was not found on the database.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

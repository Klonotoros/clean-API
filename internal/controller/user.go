package controller

import (
	"clean-API/internal/dto"
	"clean-API/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func newUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u userController) Login(context *gin.Context) {
	var loginRequest dto.LoginRequest
	err := context.ShouldBindJSON(&loginRequest)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	token, err := u.userService.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, dto.LoginResponse{Message: "Login successful!", Token: token})

}

func (u userController) Register(context *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := context.ShouldBindJSON(&registerRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	token, err := u.userService.Register(registerRequest.Email, registerRequest.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, dto.RegisterResponse{Message: "Register successful!", Token: token})
}

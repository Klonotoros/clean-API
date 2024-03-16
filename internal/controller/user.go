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
		context.JSON(http.StatusBadRequest, dto.LoginResponse{Message: "Login fail", Token: ""})
		return
	}

	token, err := u.userService.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, dto.LoginResponse{Message: "Login fail", Token: ""})
		return
	}

	context.JSON(http.StatusOK, dto.LoginResponse{Message: "Login successful!", Token: token})

}

func (u userController) Register(context *gin.Context) {
	var registerRequest dto.RegisterRequest
	if err := context.ShouldBindJSON(&registerRequest); err != nil {
		context.JSON(http.StatusBadRequest, dto.RegisterResponse{Message: "Register failed", Token: ""})
		return
	}

	token, err := u.userService.Register(registerRequest.Email, registerRequest.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto.RegisterResponse{Message: "Register failed", Token: ""})
		return
	}

	context.JSON(http.StatusCreated, dto.RegisterResponse{Message: "Register successful!", Token: token})
}

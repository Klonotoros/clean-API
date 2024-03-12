package controller

import (
	"clean-API/internal/dto"
	"clean-API/internal/middleware"
	"clean-API/internal/service"
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	User() UserController
	Info() InfoController
	Route(server *gin.Engine)
}

type controllers struct {
	userController       UserController
	infoController       InfoController
	conferenceController ConferenceController
	config               dto.Config
}

func NewControllers(services service.Services, config dto.Config) Controllers {
	userController := newUserController(services.User())
	conferenceController := newConferenceController(services.Conference())
	infoController := newInfoController()
	return &controllers{
		userController:       userController,
		conferenceController: conferenceController,
		infoController:       infoController,
		config:               config,
	}
}

func (c controllers) User() UserController {
	return c.userController
}

func (c controllers) Info() InfoController {
	return c.infoController
}

func (c controllers) Route(server *gin.Engine) {
	server.GET("/", c.infoController.Info)

	server.GET("/conferences", c.conferenceController.getConferences)
	server.GET("/conferences/:id", c.conferenceController.getConference)

	authenticated := server.Group("/")
	f := middleware.Authenticate(c.config)
	authenticated.Use(f)

	authenticated.POST("/conferences", c.conferenceController.createConference)
	authenticated.PUT("/conferences/:id", c.conferenceController.updateConference)
	authenticated.DELETE("/conferences/:id", c.conferenceController.deleteConference)

	server.POST("/signup", c.userController.Register)
	server.POST("/login", c.userController.Login)
}

package service

import (
	"clean-API/internal/client"
	"clean-API/internal/dto"
	"clean-API/internal/repository"
)

type Services interface {
	User() UserService
	Conference() ConferenceService
}

type services struct {
	userService       UserService
	conferenceService ConferenceService
}

func NewServices(repositories repository.Repositories, config dto.Config, clients client.Clients) Services {
	userService := newUserService(repositories.User(), config)
	conferenceService := newConferenceService(repositories.Conference(), config)
	return &services{
		userService:       userService,
		conferenceService: conferenceService,
	}
}

func (s services) User() UserService {
	return s.userService
}

func (s services) Conference() ConferenceService {
	return s.conferenceService
}

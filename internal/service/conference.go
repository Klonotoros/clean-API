package service

import (
	"clean-API/internal/dto"
	"clean-API/internal/model"
	"clean-API/internal/repository"
)

type ConferenceService interface {
	GetAllConferences() ([]model.Conference, error)
	GetConferenceById(id int64) (model.Conference, error)
	Save(conference model.Conference) (model.Conference, error)
	Update(conference model.Conference) error
	Delete(conference model.Conference) error
}
type conferenceService struct {
	conferenceRepository repository.ConferenceRepository
	config               dto.Config
}

func newConferenceService(conferenceRepository repository.ConferenceRepository, config dto.Config) ConferenceService {
	return &conferenceService{conferenceRepository, config}
}

func (c conferenceService) GetAllConferences() ([]model.Conference, error) {
	return c.conferenceRepository.GetAllConferences()
}

func (c conferenceService) GetConferenceById(id int64) (model.Conference, error) {
	return c.conferenceRepository.GetConferenceByID(id)
}

func (c conferenceService) Save(conference model.Conference) (model.Conference, error) {
	return c.conferenceRepository.Save(conference)
}

func (c conferenceService) Update(conference model.Conference) error {
	return c.conferenceRepository.Update(conference)
}

func (c conferenceService) Delete(conference model.Conference) error {
	return c.conferenceRepository.Delete(conference)
}

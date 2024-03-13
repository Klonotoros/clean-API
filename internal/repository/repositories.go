package repository

import "database/sql"

//go:generate mockgen -source=repositories.go -destination=repositories_mock.go -package repository

type Repositories interface {
	User() UserRepository
	Conference() ConferenceRepository
}

type repositories struct {
	userRepository       UserRepository
	conferenceRepository ConferenceRepository
}

func NewRepositories(db *sql.DB) Repositories {
	userRepository := newUserRepository(db)
	conferenceRepository := newConferenceRepository(db)

	return &repositories{
		userRepository:       userRepository,
		conferenceRepository: conferenceRepository,
	}
}

func (r repositories) User() UserRepository {
	return r.userRepository
}

func (r repositories) Conference() ConferenceRepository {
	return r.conferenceRepository
}

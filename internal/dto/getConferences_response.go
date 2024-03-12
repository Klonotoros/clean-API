package dto

import "clean-API/internal/model"

type GetConferencesResponse struct {
	Conferences []model.Conference `json:"conferences"`
}

package client

import "clean-API/internal/dto"

type Clients interface{}

type clients struct{}

func NewClients(_ dto.Config) Clients {
	return &clients{}
}

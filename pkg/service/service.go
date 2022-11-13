package service

import (
	"RestApiProj"
	"RestApiProj/pkg/repository"
)

type Poll interface {
	CreatePoll(poll RestApiProj.INCreate) (int, error)
	Polling(poll RestApiProj.INChoice) error
	GetPoll(poll RestApiProj.INGet) ([]RestApiProj.Poll, error)
}

type Service struct {
	Poll
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Poll: NewPollService(repos.Poll),
	}
}
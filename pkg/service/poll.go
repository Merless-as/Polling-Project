package service

import (
	"RestApiProj"
	"RestApiProj/pkg/repository"
)

type PollService struct {
	poll repository.Poll
}

func NewPollService(poll repository.Poll) *PollService {
	return &PollService{poll: poll}
}

func (s *PollService) CreatePoll(poll RestApiProj.INCreate) (int, error) {
	return s.poll.CreatePoll(poll)
}

func (s *PollService) Polling(poll RestApiProj.INChoice) error {
	return s.poll.Polling(poll)
}

func (s *PollService) GetPoll(poll RestApiProj.INGet) ([]RestApiProj.Poll, error) {
	return s.poll.GetPoll(poll)
}

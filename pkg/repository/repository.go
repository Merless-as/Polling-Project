package repository

import (
	"RestApiProj"
	"github.com/jmoiron/sqlx"
)

type Poll interface {
	CreatePoll(poll RestApiProj.INCreate) (int, error)
	Polling(poll RestApiProj.INChoice) error
	GetPoll(poll RestApiProj.INGet) ([]RestApiProj.Poll, error)
}

type Repository struct {
	Poll
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewPollPostgres(db),
	}
}

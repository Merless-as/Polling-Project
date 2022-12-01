package repository

import (
	"RestApiProj"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PollPostgres struct {
		db *sqlx.DB
}

func NewPollPostgres(db *sqlx.DB) *PollPostgres {
	return &PollPostgres{db: db}
}

func (r *PollPostgres) CreatePoll(poll RestApiProj.INCreate) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var pollId int
	createPollQuery := fmt.Sprint("INSERT INTO polls (name) VALUES ($1) RETURNING id")

	row := tx.QueryRow(createPollQuery, poll.Name)
	err = row.Scan(&pollId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//Тут программа должна возвращать id и продолжать работу

	slicePoll := RestApiProj.ResponseJson(poll.Choices)
	createChoiceQuery := fmt.Sprint("INSERT INTO choices (poll_id, choice_id, choice, call) VALUES ($1, $2, $3, $4)")
	for i := 0; i < len(slicePoll); i++ {
		_, err = tx.Exec(createChoiceQuery, pollId, i+1, slicePoll[i], 0)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return pollId, nil
}


func (r *PollPostgres) Polling(poll RestApiProj.INChoice) error {
	query := fmt.Sprintf("UPDATE choices SET call=call+1 WHERE poll_id=%d AND choice_id=%d", poll.PollID, poll.ChoiceID)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *PollPostgres) GetPoll(poll RestApiProj.INGet) ([]RestApiProj.Poll, error) {
	var choices []RestApiProj.Poll
	query := fmt.Sprintf("SELECT choice, call FROM choices WHERE poll_id=%d", poll.Id)
	if err := r.db.Select(&choices, query); err != nil {
		return nil, err
	}
	return choices, nil
}
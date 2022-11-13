package RestApiProj

import (
	"strings"
)

type INCreate struct {
	Name    string `json:"name"`
	Choices string `json:"choices"`
}

type OUTCreate struct {
	ID int `json:"id"`
}

type INChoice struct {
	PollID   int `json:"poll_id"`
	ChoiceID int `json:"choice_id"`
}

type INGet struct {
	Id int `json:"id"`
}

type Poll struct {
	Choice string `json:"choice" db:"choice"`
	Call   int    `json:"call" db:"call"`
}

func ResponseJson(s string) []string {
	in := strings.Trim(s, " ")
	slice := strings.Split(in, ",")
	return slice
}

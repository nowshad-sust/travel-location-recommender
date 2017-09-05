package model

type Question struct {
	Id            int      `json:"id"`
	Question      string   `json:"question"`
	Question_type string   `json:"question_type"`
	Options       []Option `json:"options"`
}

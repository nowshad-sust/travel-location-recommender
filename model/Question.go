package model

type Question struct {
	Question      string   `json:"question"`
	Question_type string   `json:"question_type"`
	Options       []Option `json:"options"`
}

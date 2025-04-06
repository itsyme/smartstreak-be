package models

type AnsweredQuestion struct {
	Question
	UserID     string `json:"user_id"`
	Correct    bool   `json:"correct"`
	Date       string `json:"date"`
	TimeTaken  int    `json:"timeTaken"`
	UserAnswer string `json:"userAnswer"`
}

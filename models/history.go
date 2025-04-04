package models

type AnsweredQuestion struct {
    ID            int    `json:"id"`
    Question      string `json:"question"`
    Correct       bool   `json:"correct"`
    Date          string `json:"date"`
    TimeTaken     int    `json:"timeTaken"`
    UserAnswer    string `json:"userAnswer"`
    CorrectAnswer string `json:"correctAnswer"`
    Explanation   string `json:"explanation"`
    Source        string `json:"source"`
}

type HistoryQuestionListItem struct {
    ID         int    `json:"id"`
    Question   string `json:"question"`
    Correct    bool   `json:"correct"`
    Date       string `json:"date"`
    TimeTaken  int    `json:"timeTaken"`
    UserAnswer string `json:"userAnswer"`
    Source     string `json:"source"`
}

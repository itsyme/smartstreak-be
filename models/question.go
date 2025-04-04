package models

type QuestionType string

const (
    MultipleChoice QuestionType = "multiple-choice"
    OpenEnded      QuestionType = "open-ended"
)

type Question struct {
    ID          int          `json:"id"`
    Type        QuestionType `json:"type"`
    Question    string       `json:"question"`
    Answer      string       `json:"answer"`
    Source      string       `json:"source"`
    Explanation *string      `json:"explanation,omitempty"`
}

type MultipleChoiceOption struct {
    Value string `json:"value"`
    Text  string `json:"text"`
}

type MultipleChoiceQuestion struct {
    Question
    Options []MultipleChoiceOption `json:"options"`
}

type OpenEndedQuestion struct {
    Question
}

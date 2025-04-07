package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/itsyme/smartstreak-be/db"
	"github.com/itsyme/smartstreak-be/models"
)

func GetQuestionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var baseQuestion models.Question
	// for multiple choice question
	var optionsJSON []byte

	err := db.DB.QueryRow(`
		SELECT id, type, question, answer, source, explanation, options
		FROM questions
		WHERE id = $1
	`, id).Scan(
		&baseQuestion.ID,
		&baseQuestion.Type,
		&baseQuestion.Question,
		&baseQuestion.Answer,
		&baseQuestion.Source,
		&baseQuestion.Explanation,
		&optionsJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("DB error: %v", err)
		}
		return
	}

	var response interface{}

	switch baseQuestion.Type {
	case models.MultipleChoice:
		var options []models.MultipleChoiceOption
		if err := json.Unmarshal(optionsJSON, &options); err != nil {
			http.Error(w, "Failed to parse options JSON", http.StatusInternalServerError)
			log.Printf("Unmarshal error: %v", err)
			return
		}
		response = models.MultipleChoiceQuestion{
			Question: baseQuestion,
			Options:  options,
		}

	case models.OpenEnded:
		response = models.OpenEndedQuestion{
			Question: baseQuestion,
		}

	default:
		response = baseQuestion
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

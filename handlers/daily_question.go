package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/itsyme/smartstreak-be/db"
	"github.com/itsyme/smartstreak-be/models"
	"github.com/itsyme/smartstreak-be/utils"
	"github.com/lib/pq"
)

const isoDateLayout = "2006-01-02"

func GetTodaysDailyQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	today := time.Now().Format(isoDateLayout)

	// Get user ID from token
	authHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := utils.GetUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Get subscription status
	var subscriptionStatus string
	err = db.DB.QueryRow("SELECT subscription_status FROM users WHERE id = $1", userID).
		Scan(&subscriptionStatus)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		log.Printf("DB error fetching user: %v", err)
		return
	}

	// Determine how many questions to fetch
	limit := 1
	if subscriptionStatus == "Scholar" {
		limit = 10
	}

	// Get limited list of question IDs directly from DB
	rows, err := db.DB.Query(`
		SELECT unnest(question_ids) AS question_id
		FROM daily_questions
		WHERE date = $1
		LIMIT $2
	`, today, limit)
	if err != nil {
		http.Error(w, "Failed to fetch daily questions", http.StatusInternalServerError)
		log.Printf("DB error retrieving question IDs: %v", err)
		return
	}
	defer rows.Close()

	var questionIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Printf("Error scanning question_id: %v", err)
			continue
		}
		questionIDs = append(questionIDs, id)
	}

	if len(questionIDs) == 0 {
		http.Error(w, "No daily questions found", http.StatusNotFound)
		return
	}

	// Fetch all questions by ID
	questionRows, err := db.DB.Query(`
		SELECT id, type, question, answer, source, explanation, options
		FROM questions
		WHERE id = ANY($1)
	`, pq.Array(questionIDs))
	if err != nil {
		http.Error(w, "Failed to fetch questions", http.StatusInternalServerError)
		log.Printf("DB error fetching questions: %v", err)
		return
	}
	defer questionRows.Close()

	// Process questions
	var questions []interface{}
	for questionRows.Next() {
		var q models.Question
		var optionsJSON []byte

		err := questionRows.Scan(
			&q.ID,
			&q.Type,
			&q.Question,
			&q.Answer,
			&q.Source,
			&q.Explanation,
			&optionsJSON,
		)
		if err != nil {
			log.Printf("Error scanning question row: %v", err)
			continue
		}

		switch q.Type {
		case models.MultipleChoice:
			var options []models.MultipleChoiceOption
			if err := json.Unmarshal(optionsJSON, &options); err != nil {
				log.Printf("Error unmarshaling options: %v", err)
				continue
			}
			questions = append(questions, models.MultipleChoiceQuestion{
				Question: q,
				Options:  options,
			})
		case models.OpenEnded:
			questions = append(questions, models.OpenEndedQuestion{
				Question: q,
			})
		default:
			questions = append(questions, q)
		}
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

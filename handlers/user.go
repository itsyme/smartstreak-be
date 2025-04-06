package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/itsyme/smartstreak-be/db"
	"github.com/itsyme/smartstreak-be/models"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var user models.User
	err := db.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.SubscriptionStatus)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("DB error: %v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

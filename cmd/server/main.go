package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/itsyme/smartstreak-be/db"
	"github.com/itsyme/smartstreak-be/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	defer db.DB.Close()

	http.HandleFunc("/user", handlers.GetUserHandler)

	http.HandleFunc("/question", handlers.GetQuestionHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go backend is running!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

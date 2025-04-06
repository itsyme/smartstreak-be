package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	DB, err = sql.Open("postgres", os.Getenv("SUPABASE_DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

}

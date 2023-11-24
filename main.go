package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func connectDB(host, port, user, pass, dbname string) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%v port=%s user=%s password=%v dbname=%v sslmode=disable", host, port, user, pass, dbname)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}

func main() {
	port := getEnv("APP_PORT", ":4444")

	db, err := connectDB(
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASS", ""),
		getEnv("DB_NAME", "postgres"),
	)

	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("db connected")
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD=%v \tURL=%v\n", r.Method, r.URL.String())
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "ok",
		})
	})

	log.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

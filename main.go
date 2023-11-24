package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sesi-11/internal/config"
	"sesi-11/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	port := config.GetEnv("APP_PORT", ":4444")

	db, err := db.ConnectDB(
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_PORT", "5432"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASS", ""),
		config.GetEnv("DB_NAME", "postgres"),
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

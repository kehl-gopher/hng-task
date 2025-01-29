package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	srv := http.NewServeMux()

	corsMiddleWare := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
	srv.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Please visit the /me endpoint for user data"))
	})
	srv.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		me := struct {
			Email           string `json:"email"`
			CurrentDateTime string `json:"current_datetime"`
			GithubUrl       string `json:"github_url"`
		}{
			Email:           "kelanidarasimi9@gmail.com",
			CurrentDateTime: time.Now().UTC().Format(time.RFC3339),
			GithubUrl:       "https://github.com/kehl-gopher/hng-task",
		}
		dt, err := json.Marshal(me)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(dt)
	})

	port := os.Getenv("PORT") // Get the port from environment variables
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}
	log.Printf("Server connected successfully on: %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), corsMiddleWare(srv))
	log.Fatal(err)
}

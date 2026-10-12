package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	port := getEnv("PORT", "8080")
	prNumber := getEnv("PR_NUMBER", "none")
	gitSHA := getEnv("GIT_SHA", "dev")

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Preview environment\nPR: %s\nSHA: %s\n", prNumber, gitSHA)
	})

	log.Printf("Listening on :%s (PR=%s), SHA=%s)", port, prNumber, gitSHA)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

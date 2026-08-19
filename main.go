package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	pool, err := openDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)
	mux.HandleFunc("POST /users", handleCreateUser(pool))
	mux.HandleFunc("GET /users/{id}", handleGetUser(pool))

	addr := ":" + env("HTTP_PORT", "8080")
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

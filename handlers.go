package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func handleCreateUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		if err := createUser(r.Context(), pool, &u); err != nil {
			log.Printf("create user: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusCreated, u)
	}
}

func handleGetUser(pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		u, err := getUser(r.Context(), pool, id)
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if err != nil {
			log.Printf("get user: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, u)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

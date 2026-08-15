package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"webproject/internal/domain"
	"webproject/internal/domain/entity"
	"webproject/internal/usecase"
)

type UserHandler struct {
	uc *usecase.UserUsecase
}

func NewUserHandler(uc *usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, domain.ErrInvalidInput.Error(), http.StatusBadRequest)
		return
	}

	if err := h.uc.CreateUser(r.Context(), &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, domain.ErrInvalidInput.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.uc.GetUser(r.Context(), id)
	if errors.Is(err, domain.ErrNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

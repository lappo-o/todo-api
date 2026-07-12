package handler

import (
	"encoding/json"
	"myapp/handler/dto"
	"myapp/service"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.UserService.Register(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("registration successful"))
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.UserService.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := dto.LoginResponse{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

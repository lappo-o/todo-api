package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"myapp/auth"
	"myapp/handler/dto"
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Service *service.TaskService
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	tasks, err := h.Service.GetAll(userID)
	if err != nil {
		fmt.Println("Ошибка в хендлере задач:", err)
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, tasks, "")
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, err)
		return
	}
	task, err := h.Service.GetByID(userID, id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, task, "")
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(w, err)
		return
	}
	newTask, err := h.Service.CreateTask(userID, req.Text)
	if err != nil {
		fmt.Println("Ошибка в хендлере задач:", err)
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, newTask, "")
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	var req dto.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(w, err)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, err)
		return
	}
	task, err := h.Service.Update(userID, id, req.Text)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, task, "")
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := auth.GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, err)
		return
	}
	err = h.Service.Delete(userID, id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "task deleted",
	}, "")
}

// Errors Handler application
type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, data any, errMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Data:  data,
		Error: errMsg,
	})
}

func mapError(err error) (int, string) {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return http.StatusNotFound, "not found"
	case errors.Is(err, model.ErrInvalid):
		return http.StatusBadRequest, "bad request"
	default:
		return http.StatusInternalServerError, "internal error"
	}
}

func handleError(w http.ResponseWriter, err error) {
	status, msg := mapError(err)
	writeJSON(w, status, nil, msg)
}

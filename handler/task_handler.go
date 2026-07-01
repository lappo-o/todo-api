package handler

import (
	"encoding/json"
	"errors"
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
	tasks, err := h.Service.GetAll()
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, tasks, "")
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, err)
		return
	}
	task, err := h.Service.GetByID(id)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, task, "")
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleError(w, err)
		return
	}
	newTask, err := h.Service.CreateTask(req.Text)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, newTask, "")
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
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
	task, err := h.Service.Update(id, req.Text)
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, task, "")
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, err)
		return
	}
	err = h.Service.Delete(id)
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

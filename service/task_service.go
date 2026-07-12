package service

import (
	"myapp/model"
	"strings"
)

type Repository interface {
	GetAll(userID int) ([]model.Task, error)
	Add(userID int, text string) (model.Task, error)
	FindByID(uesrID int, id int) (model.Task, error)
	Delete(userID int, id int) error
	Update(userID int, id int, text string) (model.Task, error)
}

type TaskService struct {
	repo Repository
}

func NewTaskService(repo Repository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(userID int, text string) (model.Task, error) {
	text = strings.TrimSpace(text)
	if len(text) < 2 || len(text) > 200 {
		return model.Task{}, model.ErrInvalid
	}
	return s.repo.Add(userID, text)
}

func (s *TaskService) GetAll(userID int) ([]model.Task, error) {
	return s.repo.GetAll(userID)
}

func (s *TaskService) GetByID(userID int, id int) (model.Task, error) {
	return s.repo.FindByID(userID, id)
}

func (s *TaskService) Delete(userID, id int) error {
	return s.repo.Delete(userID, id)
}

func (s *TaskService) Update(userID int, id int, text string) (model.Task, error) {
	if text == "" {
		return model.Task{}, model.ErrInvalid
	}
	return s.repo.Update(userID, id, text)
}

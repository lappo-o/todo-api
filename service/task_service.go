package service

import (
	"myapp/model"
	"strings"
)

type Repository interface {
	GetAll() ([]model.Task, error)
	Add(text string) (model.Task, error)
	FindByID(id int) (model.Task, error)
	Delete(id int) error
	Update(id int, text string) (model.Task, error)
}

type TaskService struct {
	repo Repository
}

func NewTaskService(repo Repository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(text string) (model.Task, error) {
	if text == "" {
		return model.Task{}, model.ErrInvalid
	}
	text = strings.TrimSpace(text)
	if len(text) > 2 && len(text) < 200 {
		return model.Task{}, model.ErrInvalid
	}
	return s.repo.Add(text)
}

func (s *TaskService) GetAll() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id int) (model.Task, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *TaskService) Update(id int, text string) (model.Task, error) {
	if text == "" {
		return model.Task{}, model.ErrInvalid
	}
	return s.repo.Update(id, text)
}

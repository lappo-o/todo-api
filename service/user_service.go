package service

import (
	"errors"
	"myapp/auth"
	"myapp/repository"
	"strings"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(email, password string) error {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	if email == "" || len(password) < 6 {
		return errors.New("invalid email or password")
	}

	return s.repo.CreateUser(email, password)
}

func (s *UserService) Login(email, password string) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("wrong email or password")
	}
	if password != user.Password {
		return "", errors.New("wrong email or password")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

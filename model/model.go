package model

import (
	"errors"
)

type Task struct {
	UserID int    `json:"user_id"`
	ID     int    `json:"id"`
	Text   string `json:"text"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrNotFound     = errors.New("not found")
	ErrInvalid      = errors.New("invalid input")
	ErrUnauthorized = errors.New("Unauthorized")
)

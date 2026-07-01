package model

import "errors"

type Task struct {
	ID   int
	Text string
}

var (
	ErrNotFound = errors.New("not found")
	ErrInvalid  = errors.New("invalid input")
)

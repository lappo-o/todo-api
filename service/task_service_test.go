package service

import (
	"myapp/model"
	"testing"
)

type FakeRepo struct{}

func (f *FakeRepo) GetAll() ([]model.Task, error) {
	return []model.Task{}, nil
}

func (f *FakeRepo) Add(text string) (model.Task, error) {
	return model.Task{
		ID:   1,
		Text: "test",
	}, nil
}

func (f *FakeRepo) FindByID(id int) (model.Task, error) {
	return model.Task{
		ID:   id,
		Text: "text",
	}, nil
}

func (f *FakeRepo) Delete(id int) error {
	return nil
}

func (f *FakeRepo) Update(id int, text string) (model.Task, error) {
	return model.Task{
		ID:   id,
		Text: text,
	}, nil
}

func TestCreateTask(t *testing.T) {
	repo := &FakeRepo{}
	service := NewTaskService(repo)
	task, err := service.CreateTask("hello")

	if err != nil {
		t.Fatal(err)
	}

	if task.Text != "hello" {
		t.Fatalf("expected hello, got %s", task.Text)
	}
}

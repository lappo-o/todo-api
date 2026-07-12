package service

// import (
// 	"errors"
// 	"myapp/model"
// 	"testing"
// )

// type FakeRepository struct {
// 	tasks         []model.Task
// 	AddCalls      int
// 	FindByIDcalls int
// }

// func (f *FakeRepository) GetAll() ([]model.Task, error) {
// 	return f.tasks, nil
// }

// func (f *FakeRepository) Add(text string) (model.Task, error) {
// 	f.AddCalls++
// 	newTask := model.Task{
// 		ID:   len(f.tasks) + 1,
// 		Text: text,
// 	}
// 	f.tasks = append(f.tasks, newTask)
// 	return newTask, nil
// }

// func (f *FakeRepository) FindByID(id int) (model.Task, error) {
// 	f.FindByIDcalls++
// 	for _, task := range f.tasks {
// 		if task.ID == id {
// 			return task, nil
// 		}
// 	}
// 	return model.Task{}, model.ErrNotFound
// }

// func (f *FakeRepository) Delete(id int) error {
// 	for i, task := range f.tasks {
// 		if task.ID == id {
// 			f.tasks = append(f.tasks[:i], f.tasks[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return model.ErrNotFound
// }

// func (f *FakeRepository) Update(id int, text string) (model.Task, error) {
// 	for i, task := range f.tasks {
// 		if task.ID == id {
// 			f.tasks[i].Text = text
// 			return f.tasks[i], nil
// 		}
// 	}
// 	return model.Task{}, model.ErrNotFound
// }

// func TestCreateTask_Success(t *testing.T) {
// 	fakeRepo := &FakeRepository{}
// 	svc := NewTaskService(fakeRepo)
// 	task, err := svc.CreateTask("Изучить тесты")

// 	if err != nil {
// 		t.Fatalf("Ожидали nil, получили ошибку: %v", err)
// 	}
// 	if task.Text != "Изучить тесты" {
// 		t.Errorf("Ожидали текст 'Изучить тесты', получили '%s'", task.Text)
// 	}
// 	if fakeRepo.AddCalls != 1 {
// 		t.Errorf("Ожидали 1 вызов Add, получили %d", fakeRepo.AddCalls)
// 	}
// }

// func TestCreateTask_TooShort(t *testing.T) {
// 	fakeRepo := &FakeRepository{}
// 	svc := NewTaskService(fakeRepo)
// 	task, err := svc.CreateTask("Х")
// 	if err == nil {
// 		t.Fatal("Ожидали ошибку, но её нет.")
// 	}
// 	if task.Text != "" {
// 		t.Errorf("Ожидали пустой текст (''), но получили '%s'", task.Text)
// 	}
// 	if !errors.Is(err, model.ErrInvalid) {
// 		t.Errorf("Ожидали '%s', получили '%s'", model.ErrInvalid, err)
// 	}
// }

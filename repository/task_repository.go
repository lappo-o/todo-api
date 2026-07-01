package repository

import (
	"database/sql"
	"myapp/model"
)

// type TaskRepository struct {
// 	tasks []model.Task
// }

// func NewTaskRepository() *TaskRepository {
// 	return &TaskRepository{
// 		tasks: []model.Task{
// 			{ID: 1, Text: "купить хлеб"},
// 			{ID: 2, Text: "выучить go"},
// 		},
// 	}
// }

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	rows, err := r.db.Query("SELECT id, text FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Text); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) Add(text string) (model.Task, error) {
	res, err := r.db.Exec(
		"INSERT INTO tasks (text) VALUES (?)",
		text,
	)
	if err != nil {
		return model.Task{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{
		ID:   int(id),
		Text: text,
	}, nil
}

func (r *TaskRepository) FindByID(id int) (model.Task, error) {
	var t model.Task

	err := r.db.QueryRow(
		"SELECT id, text FROM tasks WHERE id = ?",
		id,
	).Scan(&t.ID, &t.Text)
	if err == sql.ErrNoRows {
		return model.Task{}, model.ErrNotFound
	}

	return t, err
}

func (r *TaskRepository) Delete(id int) error {
	res, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return model.ErrNotFound
	}

	return nil
}

func (r *TaskRepository) Update(id int, text string) (model.Task, error) {
	res, err := r.db.Exec(
		"UPDATE tasks SET text = ? WHERE id = ?",
		text, id,
	)
	if err != nil {
		return model.Task{}, err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return model.Task{}, model.ErrNotFound
	}

	return model.Task{ID: id, Text: text}, nil
}

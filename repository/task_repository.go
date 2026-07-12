package repository

import (
	"database/sql"
	"myapp/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll(userID int) ([]model.Task, error) {
	rows, err := r.db.Query("SELECT user_id, id, text FROM tasks WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.UserID, &t.ID, &t.Text); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) Add(userID int, text string) (model.Task, error) {
	var id int
	err := r.db.QueryRow(
		"INSERT INTO tasks (user_id, text) VALUES ($1, $2) RETURNING id",
		userID, text,
	).Scan(&id)
	if err != nil {
		return model.Task{}, err
	}

	return model.Task{
		UserID: userID,
		ID:     int(id),
		Text:   text,
	}, nil
}

func (r *TaskRepository) FindByID(userID int, id int) (model.Task, error) {
	var t model.Task
	err := r.db.QueryRow(
		"SELECT user_id, id, text FROM tasks WHERE id = $1",
		id,
	).Scan(&t.UserID, &t.ID, &t.Text)
	if err == sql.ErrNoRows {
		return model.Task{}, model.ErrNotFound
	}

	return t, err
}

func (r *TaskRepository) Delete(userID int, id int) error {
	res, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return model.ErrNotFound
	}

	return nil
}

func (r *TaskRepository) Update(userID int, id int, text string) (model.Task, error) {
	res, err := r.db.Exec(
		"UPDATE tasks SET text = $1 WHERE id = $2",
		text, id,
	)
	if err != nil {
		return model.Task{}, err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return model.Task{}, model.ErrNotFound
	}

	return model.Task{UserID: userID, ID: id, Text: text}, nil
}

func (r *TaskRepository) Count() (int, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

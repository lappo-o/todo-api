package repository

import (
	"database/sql"
	"myapp/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(email, password string) error {
	_, err := u.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", email, password)
	return err
}

func (u *UserRepository) GetUserByEmail(email string) (model.User, error) {
	row := u.db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email)

	var user model.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

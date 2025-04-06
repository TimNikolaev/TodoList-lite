package repository

import (
	"todo-std"

	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (s *TaskPostgres) Create(userID int, item todo.Task) (int, error) {
	return 0, nil
}

func (s *TaskPostgres) GetAll(userID int) ([]todo.Task, error) {
	return nil, nil
}

func (s *TaskPostgres) GetByID(userID, taskID int) (todo.Task, error) {
	return todo.Task{}, nil
}

func (s *TaskPostgres) Delete(userID, taskID int) error {
	return nil
}

func (s *TaskPostgres) Update(userID, taskID int, input todo.Task) error {
	return nil
}

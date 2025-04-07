package repository

import (
	"fmt"
	"todo-std"

	"github.com/jmoiron/sqlx"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (s *TaskPostgres) Create(userID int, task todo.Task) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	var taskID int
	createTaskQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoTaskTable)

	if err := tx.QueryRow(createTaskQuery, task.Title, task.Description).Scan(&taskID); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersTasksQuery := fmt.Sprintf("INSERT INTO %s (user_id, task_id) VALUES ($1, $2)", usersTasksTable)
	if _, err = tx.Exec(createUsersTasksQuery, userID, taskID); err != nil {
		tx.Rollback()
		return 0, err
	}

	return taskID, tx.Commit()
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

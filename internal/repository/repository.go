package repository

import (
	"todo-std"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoTaskTable   = "todo_tasks"
	usersTasksTable = "users_tasks"
)

type Repository struct {
	todo.UserRepository
	todo.TaskRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewAuthPostgres(db),
		TaskRepository: NewTaskPostgres(db),
	}
}

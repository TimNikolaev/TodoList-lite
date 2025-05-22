package repository

import (
	"todo-std"
	"todo-std/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	todo.UserRepository
	todo.TaskRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: postgres.NewAuthPostgres(db),
		TaskRepository: postgres.NewTaskPostgres(db),
	}
}

package postgres

import (
	"todo-std/internal/config"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoTaskTable   = "todo_tasks"
	usersTasksTable = "users_tasks"
)

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.DB.DSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

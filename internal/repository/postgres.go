package repository

import (
	"todo-std/configs"

	"github.com/jmoiron/sqlx"
)

func NewPostgresDB(cfg *configs.Config) (*sqlx.DB, error) {
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

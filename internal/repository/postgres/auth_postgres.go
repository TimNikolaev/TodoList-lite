package postgres

import (
	"fmt"
	"todo-std"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var userID int

	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) values ($1, $2, $3) RETURNING id", UsersTable)

	if err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *AuthPostgres) GetUser(email, password_hash string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2 ", UsersTable)

	err := r.db.Get(&user, query, email, password_hash)

	return user, err
}

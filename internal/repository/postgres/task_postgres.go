package postgres

import (
	"fmt"
	"strings"
	"todo-std"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) Create(userID int, task todo.Task) (int, error) {
	tx, err := r.db.Begin()
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

func (r *TaskPostgres) GetAll(userID int, status string) ([]todo.Task, error) {
	var tasks []todo.Task

	var statusParam string

	switch status {
	case "true":
		statusParam = "t.done = true AND ut.user_id = $1"
	case "false":
		statusParam = "t.done = false AND ut.user_id = $1"
	default:
		statusParam = "ut.user_id = $1"
	}

	query := fmt.Sprintf("SELECT t.id, t.title, t.description, t.done FROM %s t INNER JOIN %s ut on t.id = ut.task_id WHERE %s", todoTaskTable, usersTasksTable, statusParam)

	err := r.db.Select(&tasks, query, userID)

	return tasks, err
}

func (r *TaskPostgres) GetByID(userID, taskID int) (todo.Task, error) {
	var task todo.Task

	query := fmt.Sprintf("SELECT t.id, t.title, t.description, t.done FROM %s t INNER JOIN %s ut on t.id = ut.task_id WHERE ut.user_id = $1 AND ut.task_id = $2", todoTaskTable, usersTasksTable)

	err := r.db.Get(&task, query, userID, taskID)

	return task, err
}

func (r *TaskPostgres) Update(userID, taskID int, input todo.UpdateTaskInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s t SET %s FROM %s ut WHERE t.id = ut.task_id AND ut.task_id = $%d AND ut.user_id = $%d",
		todoTaskTable,
		setQuery,
		usersTasksTable,
		argID,
		argID+1,
	)

	args = append(args, taskID, userID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TaskPostgres) Delete(userID, taskID int) error {
	query := fmt.Sprintf("DELETE FROM %s t USING %s ut WHERE t.id = ut.task_id AND ut.user_id = $1 AND ut.task_id = $2", todoTaskTable, usersTasksTable)

	_, err := r.db.Exec(query, userID, taskID)

	return err
}

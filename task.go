package todo

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskRepository interface {
	Create(userID int, item Task) (int, error)
	GetAll(userID int) ([]Task, error)
	GetByID(userID, taskID int) (Task, error)
	Delete(userID, taskID int) error
	Update(userID, taskID int, input Task) error
}

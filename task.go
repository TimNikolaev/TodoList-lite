package todo

type Task struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

type TaskRepository interface {
	Create(userID int, item Task) (int, error) //done
	GetAll(userID int) ([]Task, error)
	GetByID(userID, taskID int) (Task, error)
	Delete(userID, taskID int) error
	Update(userID, taskID int, input UpdateTaskInput) error
}

package todo

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRepository interface {
	CreateUser(user User) (int, error)
	GetUser(email, password string) (User, error)
}

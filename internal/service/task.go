package service

import "todo-std"

func (s *Service) Create(userID int, item todo.Task) (int, error) {
	return 0, nil
}

func (s *Service) GetAll(userID int) ([]todo.Task, error) {
	return nil, nil
}

func (s *Service) GetByID(userID, taskID int) (todo.Task, error) {
	return todo.Task{}, nil
}

func (s *Service) Delete(userID, taskID int) error {
	return nil
}

func (s *Service) Update(userID, taskID int, input todo.Task) error {
	return nil
}

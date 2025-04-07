package service

import "todo-std"

func (s *Service) CreateTask(userID int, item todo.Task) (int, error) {
	return s.taskRepository.Create(userID, item)
}

func (s *Service) GetAllTasks(userID int) ([]todo.Task, error) {
	return nil, nil
}

func (s *Service) GetTaskByID(userID, taskID int) (todo.Task, error) {
	return todo.Task{}, nil
}

func (s *Service) DeleteTask(userID, taskID int) error {
	return nil
}

func (s *Service) UpdateTask(userID, taskID int, input todo.Task) error {
	return nil
}

package service

import "todo-std"

func (s *Service) CreateTask(userID int, item todo.Task) (int, error) {
	return s.TaskRepository.Create(userID, item)
}

func (s *Service) GetAllTasks(userID int, status string) ([]todo.Task, error) {
	return s.TaskRepository.GetAll(userID, status)
}

func (s *Service) GetTaskByID(userID, taskID int) (todo.Task, error) {
	return s.TaskRepository.GetByID(userID, taskID)
}

func (s *Service) UpdateTask(userID, taskID int, input todo.UpdateTaskInput) error {
	return s.TaskRepository.Update(userID, taskID, input)
}

func (s *Service) DeleteTask(userID, taskID int) error {
	return s.TaskRepository.Delete(userID, taskID)
}

package service

import (
	"todo-std"
	"todo-std/configs"
)

type Service struct {
	taskRepository todo.TaskRepository
	userRepository todo.UserRepository
	Config         *configs.Config
}

func NewService(taskRepo todo.TaskRepository, userRepo todo.UserRepository, config *configs.Config) *Service {
	return &Service{
		taskRepository: taskRepo,
		userRepository: userRepo,
		Config:         config,
	}
}

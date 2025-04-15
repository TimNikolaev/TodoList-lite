package service

import (
	"todo-std"
	"todo-std/configs"
	"todo-std/internal/repository"
)

type Service struct {
	taskRepository todo.TaskRepository
	userRepository todo.UserRepository
	Config         *configs.Config
}

func NewService(repo *repository.Repository, config *configs.Config) *Service {
	return &Service{
		taskRepository: repo.TaskRepository,
		userRepository: repo.UserRepository,
		Config:         config,
	}
}

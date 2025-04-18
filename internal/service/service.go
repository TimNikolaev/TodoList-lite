package service

import (
	"todo-std"
	"todo-std/configs"
	"todo-std/internal/repository"
)

type Service struct {
	todo.TaskRepository
	todo.UserRepository
	*configs.Config
}

func NewService(repo *repository.Repository, config *configs.Config) *Service {
	return &Service{
		TaskRepository: repo.TaskRepository,
		UserRepository: repo.UserRepository,
		Config:         config,
	}
}

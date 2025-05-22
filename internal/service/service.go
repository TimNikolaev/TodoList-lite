package service

import (
	"todo-std"
	"todo-std/internal/config"
	"todo-std/internal/repository"
)

type Service struct {
	todo.TaskRepository
	todo.UserRepository
	*config.Config
}

func NewService(repo *repository.Repository, config *config.Config) *Service {
	return &Service{
		TaskRepository: repo.TaskRepository,
		UserRepository: repo.UserRepository,
		Config:         config,
	}
}

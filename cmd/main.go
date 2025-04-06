package main

import (
	"todo-std"
	"todo-std/configs"
	"todo-std/internal/handler"
	"todo-std/internal/repository"
	"todo-std/internal/service"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	config := configs.LoadConfig()
	db, err := repository.NewPostgresDB(config)
	if err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository.TaskRepository, repository.UserRepository, config)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

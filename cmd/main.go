package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"todo-std"
	"todo-std/internal/config"
	"todo-std/internal/handler"
	"todo-std/internal/repository"
	"todo-std/internal/repository/postgres"
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

	config := config.LoadConfig()
	db, err := postgres.NewPostgresDB(config)
	if err != nil {
		logrus.Fatalf("error initialization configs: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository, config)
	handler := handler.NewHandler(service)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRouts()); err != nil {
			logrus.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shotting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

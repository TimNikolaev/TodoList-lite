package handler

import (
	"todo-std/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	task := router.Group("/task", h.userIdentity)
	{
		task.POST("/", h.createTask)
		task.GET("/:id", h.getTaskById)
		task.GET("/", h.getAllTasks)
		task.PUT("/:id", h.updateTask)
		task.DELETE("/:id", h.deleteTask)
	}

	return router
}

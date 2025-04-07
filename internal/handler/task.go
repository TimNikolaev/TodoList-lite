package handler

import (
	"net/http"
	"todo-std"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input todo.Task

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	taskID, err := h.service.CreateTask(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{"id": taskID})

}

func (h *Handler) getAllTasks(c *gin.Context) {

}

func (h *Handler) getTaskById(c *gin.Context) {

}

func (h *Handler) updateTask(c *gin.Context) {

}

func (h *Handler) deleteTask(c *gin.Context) {

}

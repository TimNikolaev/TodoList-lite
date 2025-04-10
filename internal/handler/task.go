package handler

import (
	"net/http"
	"strconv"
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

type getAllTasksResponse struct {
	Data []todo.Task `json:"data"`
}

func (h *Handler) getAllTasks(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	status := c.Query("status")

	tasks, err := h.service.GetAllTasks(userID, status)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTasksResponse{Data: tasks})

}

func (h *Handler) getTaskById(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	task, err := h.service.GetTaskByID(userID, taskID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) updateTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdateTask(userID, taskID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

func (h *Handler) deleteTask(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.DeleteTask(userID, taskID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

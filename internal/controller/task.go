package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/brentellingson/learning-golang-api/internal/dto"
	"github.com/gin-gonic/gin"
)

// TaskInterface defines the methods required for the Task service.
type TaskInterface interface {
	GetTasks(c context.Context) ([]dto.Task, error)
	GetTask(c context.Context, id int) (dto.Task, error)
	InsertTask(c context.Context, task dto.InsertTask) (dto.Task, error)
	UpdateTask(c context.Context, id int, task dto.UpdateTask) (dto.Task, error)
	DeleteTask(c context.Context, id int) error
}

// TaskController provides the implementation of the task endpoints
type TaskController struct {
	svc TaskInterface
}

// NewTaskController creates a new TaskController.
func NewTaskController(svc TaskInterface) *TaskController {
	return &TaskController{svc}
}

// AddRoutes adds the task routes to the router.
func (ctl *TaskController) AddRoutes(r *gin.RouterGroup) {
	r.GET("/tasks", ctl.GetTasks)
	r.GET("/tasks/:id", ctl.GetTask)
	r.POST("/tasks", ctl.InsertTask)
	r.PATCH("/tasks/:id", ctl.UpdateTask)
	r.DELETE("/tasks/:id", ctl.DeleteTask)
}

// GetTasks godoc
//
//	@Summary		get a list of tasks
//	@Description	get a list of tasks
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.Task
//	@Router			/tasks [get]
func (ctl *TaskController) GetTasks(c *gin.Context) {
	tasks, err := ctl.svc.GetTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask godoc
//
//	@Summary		get a task
//	@Description	get a task
//	@Tags			tasks
//	@Accept			"*/*"
//	@Produce		json
//	@Param			id	path		int	true	"task ID"
//	@Success		200	{object}	dto.Task
//	@Router			/tasks/{id} [get]
func (ctl *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	task, err := ctl.svc.GetTask(c, idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// InsertTask godoc
//
//	@Summary		insert a task
//	@Description	insert a task
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		dto.InsertTask	true	"task"
//	@Success		201		{object}	dto.Task
//	@Router			/tasks [post]
func (ctl *TaskController) InsertTask(c *gin.Context) {
	var task dto.InsertTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := ctl.svc.InsertTask(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask godoc
//
//	@Summary		update a task
//	@Description	update a task
//	@Tags			tasks
//	@Accept			json
//	@Product		json
//	@Param			id		path		int				true	"task ID"
//	@Param			task	body		dto.UpdateTask	true	"task"
//	@Success		200		{object}	dto.Task
//	@Router			/tasks/{id} [patch]
func (ctl *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var task dto.UpdateTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := ctl.svc.UpdateTask(c, idx, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask godoc
//
//	@Summary		delete a task
//	@Description	delete a task
//	@Tags			tasks
//	@Accept			"*/*"
//	@Param			id	path	int	true	"task ID"
//	@Success		204
//	@Router			/tasks/{id} [delete]
func (ctl *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	if err := ctl.svc.DeleteTask(c, idx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

package service

import (
	"context"

	"github.com/brentellingson/learning-golang-api/internal/dto"
)

// TaskService provides the business logic for tasks.
type TaskService struct{}

// NewTaskService creates a new TaskService.
func NewTaskService() *TaskService {
	return &TaskService{}
}

// GetTasks returns a list of all tasks.
func (s *TaskService) GetTasks(_ context.Context) ([]dto.Task, error) {
	return nil, nil
}

// GetTask returns a task by ID.
func (s *TaskService) GetTask(_ context.Context, _ int) (dto.Task, error) {
	return dto.Task{}, nil
}

// InsertTask inserts a new task.
func (s *TaskService) InsertTask(_ context.Context, _ dto.InsertTask) (dto.Task, error) {
	return dto.Task{}, nil
}

// UpdateTask updates a task by ID.
func (s *TaskService) UpdateTask(_ context.Context, _ int, _ dto.UpdateTask) (dto.Task, error) {
	return dto.Task{}, nil
}

// DeleteTask deletes a task by ID.
func (s *TaskService) DeleteTask(_ context.Context, _ int) error {
	return nil
}

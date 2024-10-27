package service

import (
	"context"
	"testing"

	"github.com/brentellingson/learning-golang-api/internal/dto"
	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestGetTasks(t *testing.T) {
	ts := NewTaskService()
	_, err := ts.GetTasks(context.Background())
	assert.NoError(t, err)
}

func TestGetTask(t *testing.T) {
	ts := NewTaskService()
	task, err := ts.GetTask(context.Background(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func TestInsertTask(t *testing.T) {
	ts := NewTaskService()
	newTask := dto.InsertTask{
		Title: "New Task",
	}
	task, err := ts.InsertTask(context.Background(), newTask)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func TestUpdateTask(t *testing.T) {
	ts := NewTaskService()
	updateTask := dto.UpdateTask{
		Title: ptr("Updated Task"),
	}
	task, err := ts.UpdateTask(context.Background(), 1, updateTask)
	assert.NoError(t, err)
	assert.NotNil(t, task)
}

func TestDeleteTask(t *testing.T) {
	ts := NewTaskService()
	err := ts.DeleteTask(context.Background(), 1)
	assert.NoError(t, err)
}

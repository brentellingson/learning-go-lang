package dto

// Task represents a to-do task.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
}

// InsertTask represents a request to insert a task.
type InsertTask struct {
	Title string `json:"title" binding:"required"`
}

// UpdateTask represents a request to update a task.
type UpdateTask struct {
	Title *string `json:"title,omitempty"`
}

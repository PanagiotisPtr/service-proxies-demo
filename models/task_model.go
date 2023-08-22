package models

type TaskStatus string

const (
	Todo  TaskStatus = "todo"
	Doing TaskStatus = "doing"
	Done  TaskStatus = "done"
)

type Task struct {
	Id          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}

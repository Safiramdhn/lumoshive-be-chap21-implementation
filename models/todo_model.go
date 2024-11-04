package models

type Todos struct {
	ID                  int    `json:"id,omitempty"`
	UserID              int    `json:"user_id,omitempty"`
	Description         string `json:"description,omitempty"`
	TodoStatus          string `json:"todo_status,omitempty"`
	TotalTodoNotStarted int    `json:"total_todo_not_started,omitempty"`
	TotalTodoDone       int    `json:"total_todo_done,omitempty"`
}

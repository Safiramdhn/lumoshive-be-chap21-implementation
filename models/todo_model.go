package models

type Todos struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	TodoStatus  string `json:"todo_status"`
}

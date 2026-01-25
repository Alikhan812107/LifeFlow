package models

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	UserID int    `json:"user_id"`
}

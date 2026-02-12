package models

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	TasksNum int    `json:"tasks_num" bson:"tasks_num"`
	NotesNum int    `json:"notes_num" bson:"notes_num"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Role     string `json:"role" bson:"role"`
}

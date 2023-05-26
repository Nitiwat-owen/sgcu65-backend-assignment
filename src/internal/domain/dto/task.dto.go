package dto

type TaskDto struct {
	ID       string    `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	Status   string    `json:"status" validate:"required"`
	Deadline string    `json:"deadline" validate:"required"`
	Users    []UserDto `json:"users"`
}

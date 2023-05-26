package dto

type UserDto struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" validate:"required"`
	Firstname string    `json:"firstname" validate:"required"`
	Surname   string    `json:"surname" validate:"required"`
	Role      string    `json:"role" validate:"required"`
	Position  string    `json:"position" validate:"required"`
	Salary    int       `json:"salary" validate:"required"`
	Tasks     []TaskDto `json:"tasks"`
}

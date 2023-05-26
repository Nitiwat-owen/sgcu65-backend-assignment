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

type FindUsersQueryParams struct {
	Firstname string `query:"firstname"`
	Surname   string `query:"surname"`
	Position  string `query:"position"`
}

type UpdateUserDto struct {
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Position  string `json:"position"`
	Salary    int    `json:"salary"`
}

package dto

type TaskDto struct {
	ID       string    `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	Status   string    `json:"status" validate:"required"`
	Deadline string    `json:"deadline" validate:"required"`
	Users    []UserDto `json:"users"`
}

type FindTaskQueryParams struct {
	Name string `query:"name"`
}

type UpdateTaskDto struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Status   string `json:"status"`
	Deadline string `json:"deadline"`
}

type AssignTaskDto struct {
	TaskId string `json:"taskId" validate:"uuid"`
	UserId string `json:"userId" validate:"uuid"`
}

package entity

import (
	"sgcu65-backend-assignment/src/internal/domain/dto"
	"time"
)

type Task struct {
	Base
	Name     string    `json:"name" gorm:"index"`
	Content  string    `json:"content"`
	Status   string    `json:"status"`
	Deadline time.Time `json:"deadline" gorm:"type:timestamp"`
	Users    []*User   `gorm:"many2many:user_tasks"`
}

func (task *Task) EntityToDto() *dto.TaskDto {
	result := &dto.TaskDto{
		ID:       task.ID.String(),
		Name:     task.Name,
		Content:  task.Content,
		Status:   task.Status,
		Deadline: task.Deadline.String(),
		Users:    nil,
	}

	if task.Users != nil {
		var users []dto.UserDto
		for _, user := range task.Users {
			users = append(users, *user.EntityToDto())
		}
		result.Users = users
	}

	return result
}

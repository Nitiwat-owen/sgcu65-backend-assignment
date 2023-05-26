package entity

import "sgcu65-backend-assignment/src/internal/domain/dto"

type User struct {
	Base
	Email     string `json:"email" gorm:"unique"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	//Password  string  `json:"password"`
	Role     string  `json:"role"`
	Position string  `json:"position"`
	Salary   int     `json:"salary"`
	Tasks    []*Task `gorm:"many2many:user_tasks"`
}

func (user *User) EntityToDto() *dto.UserDto {
	result := &dto.UserDto{
		ID:        user.ID.String(),
		Email:     user.Email,
		Firstname: user.Firstname,
		Surname:   user.Surname,
		Role:      user.Role,
		Position:  user.Position,
		Salary:    user.Salary,
		Tasks:     nil,
	}

	if user.Tasks != nil {
		var tasks []dto.TaskDto
		for _, task := range user.Tasks {
			tasks = append(tasks, *task.EntityToDto())
		}
		result.Tasks = tasks
	}

	return result
}

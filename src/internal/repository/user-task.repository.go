package repository

import (
	"gorm.io/gorm"
	"sgcu65-backend-assignment/src/internal/domain/entity"
)

type IUserTaskRepository interface {
	AssignTask(userTask *entity.UserTask) error
	RemoveUser(userId string, taskId string) error
}

type UserTaskRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserTaskRepositoryImpl) AssignTask(userTask *entity.UserTask) error {
	return r.DB.Create(userTask).Error
}

func (r *UserTaskRepositoryImpl) RemoveUser(userId string, taskId string) error {
	return r.DB.Where("user_id = ? AND task_id = ?", userId, taskId).Delete(&entity.UserTask{}).Error
}

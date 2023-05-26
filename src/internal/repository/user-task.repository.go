package repository

import (
	"gorm.io/gorm"
	"sgcu65-backend-assignment/src/internal/domain/entity"
)

type IUserTaskRepository interface {
	AssignTask(userTask *entity.UserTask) error
	RemoveUser(userTask *entity.UserTask) error
}

type UserTaskRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserTaskRepositoryImpl) AssignTask(userTask *entity.UserTask) error {
	return r.DB.Create(userTask).Error
}

func (r *UserTaskRepositoryImpl) RemoveUser(userTask *entity.UserTask) error {
	return r.DB.Where("user_id = ? AND task_id = ?", userTask.UserID, userTask.TaskID).Delete(&entity.UserTask{}).Error
}

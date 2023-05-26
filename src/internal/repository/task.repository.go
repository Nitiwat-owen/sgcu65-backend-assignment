package repository

import (
	"gorm.io/gorm"
	"sgcu65-backend-assignment/src/internal/domain/entity"
)

type ITaskRepository interface {
	CreateTask(task *entity.Task) error
	FindAllTasks(tasks *[]*entity.Task) error
	FindTaskByName(tasks *[]*entity.Task, name string) error
	FindTaskById(id string, task *entity.Task) error
	UpdateTask(id string, task *entity.Task) error
	DeleteTask(id string) error
}

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func (r *TaskRepositoryImpl) CreateTask(task *entity.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepositoryImpl) FindAllTasks(tasks *[]*entity.Task) error {
	return r.DB.Find(tasks).Error
}

func (r *TaskRepositoryImpl) FindTaskByName(tasks *[]*entity.Task, name string) error {
	return r.DB.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").Find(tasks).Error
}

func (r *TaskRepositoryImpl) FindTaskById(id string, task *entity.Task) error {
	return r.DB.Where("id = ?", id).First(task).Error
}

func (r *TaskRepositoryImpl) UpdateTask(id string, task *entity.Task) error {
	return r.DB.Where("id = ?", id).Updates(task).Error
}

func (r *TaskRepositoryImpl) DeleteTask(id string) error {
	return r.DB.Where("id = ?", id).Delete(&entity.Task{}).Error
}

package repository

import (
	"gorm.io/gorm"
	"sgcu65-backend-assignment/src/internal/domain/entity"
)

type IUserRepository interface {
	CreateUser(user *entity.User) error
	FindAllUsers(users *[]*entity.User) error
	FindUserByKeyword(users *[]*entity.User, firstname string, surname string, position string) error
	FindUserById(id string, user *entity.User) error
	UpdateUser(id string, user *entity.User) error
	DeleteUser(id string) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) CreateUser(user *entity.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepositoryImpl) FindAllUsers(users *[]*entity.User) error {
	return r.DB.Find(users).Error
}

func (r *UserRepositoryImpl) FindUserByKeyword(users *[]*entity.User, firstname string, surname string, position string) error {
	return r.DB.Where("LOWER(firstname) LIKE LOWER(?) OR LOWER(surname) LIKE LOWER(?) OR LOWER(position) LIKE LOWER(?)", "%"+firstname+"%", "%"+surname+"%", "%"+position+"%").Find(users).Error
}

func (r *UserRepositoryImpl) FindUserById(id string, user *entity.User) error {
	return r.DB.Preload("Tasks").Where("id = ?", id).First(user).Error
}

func (r *UserRepositoryImpl) UpdateUser(id string, user *entity.User) error {
	return r.DB.Where("id = ?", id).Updates(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id string) error {
	return r.DB.Where("id = ?", id).Delete(&entity.User{}).Error
}

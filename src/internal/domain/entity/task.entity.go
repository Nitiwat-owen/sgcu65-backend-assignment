package entity

import "time"

type Task struct {
	Base
	Name     string    `json:"name" gorm:"index"`
	Content  string    `json:"content"`
	status   string    `json:"status"`
	Deadline time.Time `json:"deadline" gorm:"type:timestamp"`
	Users    []*User   `gorm:"many2many:user_tasks"`
}

package entity

type Task struct {
	Base
	Name    string `json:"name" gorm:"index"`
	Content string `json:"content"`
	status  string `json:"status"`
}

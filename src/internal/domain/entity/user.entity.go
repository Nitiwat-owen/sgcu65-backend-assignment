package entity

type User struct {
	Base
	Email     string  `json:"email" gorm:"unique"`
	Firstname string  `json:"firstname"`
	Surname   string  `json:"surname"`
	Password  string  `json:"password"`
	Role      string  `json:"role"`
	Position  string  `json:"position"`
	Salary    int     `json:"salary"`
	Tasks     []*Task `gorm:"many2many:user_tasks"`
}

package entity

type User struct {
	Base
	Email     string `json:"email" gorm:"unique"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Role      string `json:"role"`
	Salary    int    `json:"salary"`
}

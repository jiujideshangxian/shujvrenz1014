package models

type User struct {
	Id int `from:"id"`
	Phone string `from:"phone"`
	Password string `from:"password"`
}

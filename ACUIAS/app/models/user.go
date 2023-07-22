package models

type User struct {
	Id       int    `gorm:"primary;column:id"`
	Username string `gorm:"not null;column:username"`
	Password string `gorm:"not null;column:password"`
	Phone    string `gorm:"not null;column:phone"`
	Email    string `gorm:"not null;column:email"`
	Token    string `gorm:"not null;column:token"`
}

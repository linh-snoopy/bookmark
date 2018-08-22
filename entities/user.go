package entities

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	InsertUser(email, password string) error
	UpdatePassword(password string) error
	DeleteUser(id int) error 
}

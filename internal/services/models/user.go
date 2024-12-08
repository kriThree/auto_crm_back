package models

import (
	"time"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	Roles     UserRoles
}

type UserRoles struct {
	Admin int64
	Owner int64
	Client int64
}

type RegisterUserDto struct {
	Name     string
	Email    string
	Password string
	Role     string
}

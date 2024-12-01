package storage_models

import "time"

type User struct {
	Id        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type AddUserDto struct {
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}

type UpdateUserDto struct {
	Name     string
	Email    string
	Password string
}

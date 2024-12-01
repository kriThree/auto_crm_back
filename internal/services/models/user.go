package models

import "time"

type User struct {
	Id         int64
	Name       string
	Email      string
	Password   string
	CreatedAt time.Time
}

package rest_models

import "time"

type User struct { 
	Id        int64 `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterReqDto struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type RegisterResDto struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
type UserUpdateReqDto struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginReqDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResDto struct {
	Token string `json:"token"`
	User  User `json:"user"`
}
type AuthorizeResDto struct {
	Token string `json:"token"`
	User  User `json:"user"`
}
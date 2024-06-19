package model

type User struct {
    Uid string
    Name string
}

func NewUser() *User {
    return &User{}
}

func (user *User) IsEmpty() bool {
    return User{} == *user
}

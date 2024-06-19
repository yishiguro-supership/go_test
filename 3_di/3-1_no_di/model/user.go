package model

type User struct {
    Uid string
    Name string
}

func NewUser() *User {
    return &User{}
}

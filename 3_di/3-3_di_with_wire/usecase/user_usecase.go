package usecase

// usecase層：ビジネスロジックやアプリケーションの主要機能の実現

import (
    "myproject/model"
    "myproject/repository"
)

type UserUsecase interface {
    GetUser(user *model.User, uid string)
    RegisterUser(user *model.User)
}

type userUsecase struct {
    ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
    return &userUsecase{ur}
}

func (uu *userUsecase) GetUser(user *model.User, uid string) {
    uu.ur.GetUser(user, uid)
}

func (uu *userUsecase) RegisterUser(user *model.User) {
    uu.ur.SetUser(user)
}

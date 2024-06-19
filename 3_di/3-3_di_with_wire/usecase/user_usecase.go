package usecase

// usecase層：ビジネスロジックやアプリケーションの主要機能の実現

import (
    "myproject/model"
    "myproject/repository"
)

type UserUsecase interface {
    GetUser(user *model.User, uid string)
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

package usecase

// usecase層：ビジネスロジックやアプリケーションの主要機能の実現

import (
    "errors"
    "myproject/model"
    "myproject/repository"
)

type UserUsecase interface {
    GetUser(user *model.User, uid string)
    RegisterUser(user *model.User) error
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

func (uu *userUsecase) RegisterUser(user *model.User) error {
    if uu.ur.UserExists(user.Uid) {
        return errors.New("uid already in use")
    }
    uu.ur.SetUser(user)
    return nil
}

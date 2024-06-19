package usecase

// usecase層：ビジネスロジックやアプリケーションの主要機能の実現

import (
    "myproject/repository"
)

type UserUsecase struct {
    Ur *repository.UserRepository // NG：外部パッケージから利用しないといけないため、大文字で定義
}

func NewUserUsecase() *UserUsecase {
    ur := repository.NewUserRepository() // NG：UserRepositoryに依存している
    return &UserUsecase{ur}
}

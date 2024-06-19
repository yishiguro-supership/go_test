package repository

// repository層：外部データの読み書き

import (
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "myproject/infra"
)

type UserRepository struct {
    Db *dynamodb.DynamoDB // NG：外部パッケージから利用しないといけないため、大文字で定義
}

func NewUserRepository() *UserRepository {
    db := infra.InitDB() // NG：DBに依存している
    return &UserRepository{db}
}

package repository

// repository層：外部データの読み書き

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "myproject/model"
)

type UserRepository interface {
    GetUser(user *model.User, uid string)
    SetUser(user *model.User)
}

type userRepository struct {
    db *dynamodb.DynamoDB
}

func NewUserRepository(db *dynamodb.DynamoDB) UserRepository {
    return &userRepository{db}
}

func (ur *userRepository) GetUser(user *model.User, uid string) {
    db := ur.db
    input := &dynamodb.GetItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "uid": {
                S: aws.String(uid),
            },
        },
        TableName: aws.String("user"),
    }
    result, _ := db.GetItem(input)

    if result.Item != nil {
        user.Uid = *result.Item["uid"].S
        user.Name = *result.Item["name"].S
    }
}

func (ur *userRepository) SetUser(user *model.User) {
    db := ur.db
    input := &dynamodb.PutItemInput{
        Item: map[string]*dynamodb.AttributeValue{
            "uid": {
                S: aws.String(user.Uid),
            },
            "name": {
                S: aws.String(user.Name),
            },
        },
        TableName: aws.String("user"),
    }
    db.PutItem(input)
}

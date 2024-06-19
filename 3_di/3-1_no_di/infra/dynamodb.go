package infra

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func InitDB() *dynamodb.DynamoDB {
    session, _ := session.NewSessionWithOptions(session.Options{
        Profile: "dummy",
        Config: aws.Config{
            Region: aws.String("ap-northeast-1"),
            Endpoint: aws.String("http://localhost:8000"),
        },
    })
    return dynamodb.New(session)
}

package controller

// controller層：ユーザの入力を解釈し、usecaseに伝える層

import (
    "net/http"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/gin-gonic/gin"
    "myproject/model"
    "myproject/usecase"
)

type UserController struct {
    Uu *usecase.UserUsecase // NG：外部パッケージから利用しないといけないため、大文字で定義
}

func NewUserController() *UserController {
    uu := usecase.NewUserUsecase() // NG：UserUsecaseに依存している
    return &UserController{uu}
}

func (uc *UserController) ShowUserInfo(c *gin.Context) {
    uid := c.Param("uid")
    uu := *(*uc).Uu // NG：UserControllerがUuを持っていることを知っている前提になっている
    ur := *uu.Ur // NG：UserUsecaseがUrを持っていることを知っている前提になっている
    db := ur.Db // NG：UserRepositoryがDbを持っていることを知っている前提になっている
    input := &dynamodb.GetItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "uid": {
                S: aws.String(uid),
            },
        },
        TableName: aws.String("user"),
    }
    result, _ := db.GetItem(input)
    user := model.NewUser()
    if result.Item != nil {
        user.Uid = *result.Item["uid"].S
        user.Name = *result.Item["name"].S

        c.JSON(http.StatusOK, gin.H{
            "uid": user.Uid,
            "name": user.Name,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{})
    }
}

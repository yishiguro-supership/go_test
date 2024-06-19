package controller

// controller層：ユーザの入力を解釈し、usecaseに伝える層

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "myproject/model"
    "myproject/usecase"
)

type UserController interface {
    ShowUserInfo(c *gin.Context)
}

type userController struct {
    uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
    return &userController{uu}
}

func (uc *userController) ShowUserInfo(c *gin.Context) {
    uid := c.Param("uid")
    user := model.NewUser()
    uc.uu.GetUser(user, uid)

    if !user.IsEmpty() {
        c.JSON(http.StatusOK, gin.H{
            "uid": user.Uid,
            "name": user.Name,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{})
    }
}

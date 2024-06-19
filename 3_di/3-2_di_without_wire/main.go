package main

// 呼び出し層：全体を操作する層

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "myproject/controller"
    "myproject/infra"
    "myproject/repository"
    "myproject/usecase"
)

func main() {
    db := infra.InitDB()
    ur := repository.NewUserRepository(db)
    uu := usecase.NewUserUsecase(ur)
    uc := controller.NewUserController(uu)

    engine := gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "for health-check")
    })
    engine.GET("/user/:uid", func(c *gin.Context) { uc.ShowUserInfo(c) })

    engine.Run(":3000")
}

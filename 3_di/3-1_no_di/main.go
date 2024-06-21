package main

// 呼び出し層：全体を操作する層

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "myproject/controller"
)

func main() {
    uc := controller.NewUserController()
    engine := gin.Default()

    engine.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "for health-check")
    })
    engine.GET("/user/:uid", func(c *gin.Context) { uc.ShowUserInfo(c) })
    engine.GET("/register", func(c *gin.Context) { uc.RegisterUser(c) })

    engine.Run(":3000")
}

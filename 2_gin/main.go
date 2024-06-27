package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    engine := gin.Default()

    // 実行する関数を直接記載
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello, world!",
        })
    })

    // 別で記載した関数を呼び出し
    engine.GET("/hello", func(c *gin.Context) { HelloWithoutParam(c) })
    engine.GET("/hello/:name", func(c *gin.Context) { HelloWithParam(c) })

    engine.Run(":3000")
}

// クエリを利用
func HelloWithoutParam(c *gin.Context) {
    name := c.DefaultQuery("name", "NoName")
    message := c.Query("message")
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, " + name + "!" + message,
    })
}

// パスの一部分をパラメータとして利用
func HelloWithParam(c *gin.Context) {
    name := c.Param("name")
    message := c.Query("message")
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, " + name + "!" + message,
    })
}

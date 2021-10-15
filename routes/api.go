package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    router.GET("/test", func(c *gin.Context) {
        time.Sleep(5*time.Second)
        c.String(http.StatusOK, "success")
    })
}

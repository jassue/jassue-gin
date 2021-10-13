package main

import (
    "github.com/gin-gonic/gin"
    "jassue-gin/bootstrap"
    "jassue-gin/global"
    "net/http"
)

func main() {
    // 初始化配置
    bootstrap.InitializeConfig()

    // 初始化日志
    global.App.Log = bootstrap.InitializeLog()
    global.App.Log.Info("log init success!")

    // 初始化数据库
    global.App.DB = bootstrap.InitializeDB()
    // 程序关闭前，释放数据库连接
    defer func() {
        if global.App.DB != nil {
            db, _ := global.App.DB.DB()
            db.Close()
        }
    }()

    r := gin.Default()

    // 测试路由
    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    // 启动服务器
    r.Run(":" + global.App.Config.App.Port)
}

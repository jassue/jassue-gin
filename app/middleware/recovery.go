package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/jassue/jassue-gin/app/common/response"
    "github.com/jassue/jassue-gin/global"
    "gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
    return gin.RecoveryWithWriter(
        &lumberjack.Logger{
            Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Log.Filename,
            MaxSize:    global.App.Config.Log.MaxSize,
            MaxBackups: global.App.Config.Log.MaxBackups,
            MaxAge:     global.App.Config.Log.MaxAge,
            Compress:   global.App.Config.Log.Compress,
        },
        response.ServerError)
}

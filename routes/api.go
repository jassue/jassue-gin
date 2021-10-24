package routes

import (
    "github.com/gin-gonic/gin"
    "jassue-gin/app/controllers/app"
    "jassue-gin/app/middleware"
)

// SetApiGroupRoutes api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
    router.POST("/auth/register", app.Register)
    router.POST("/auth/login", app.Login)

    authRouter := router.Group("").Use(middleware.JWTAuth(app.GuardName))
    {
        authRouter.POST("/auth/info", app.Info)
    }
}

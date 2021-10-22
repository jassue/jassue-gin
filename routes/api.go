package routes

import (
    "github.com/gin-gonic/gin"
    "jassue-gin/app/controllers/app"
)

// SetApiGroupRoutes api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
    router.POST("/user/register", app.Register)
}

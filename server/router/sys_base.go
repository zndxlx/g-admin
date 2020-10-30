package router

import (
    "github.com/gin-gonic/gin"
    "g-admin/middleware"
    "g-admin/api"
)

func InitBaseRouter(Router *gin.RouterGroup) {
    BaseRouter := Router.Group("base").Use(middleware.GinDetailLogger())
    {
        BaseRouter.POST("register", api.Register)
        BaseRouter.POST("login", api.Login)
        BaseRouter.POST("captcha", api.Captcha)
    }
    return
}

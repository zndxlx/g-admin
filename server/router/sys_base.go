package router

import (
    "github.com/gin-gonic/gin"
    "g-admin/api/base"
    "g-admin/middleware"
)

func InitBaseRouter(Router *gin.RouterGroup) {
    BaseRouter := Router.Group("base").Use(middleware.GinDetailLogger())
    {
       // BaseRouter.POST("register", v1.Register)
        BaseRouter.POST("login", base.Login)
        BaseRouter.POST("captcha", base.Captcha)
    }
    return
}

package router

import (
    // "g-admin/middleware"
    "github.com/gin-gonic/gin"
    "g-admin/api"
    "g-admin/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
    UserRouter := Router.Group("user").
        Use(middleware.GinDetailLogger()).
        Use(middleware.JWTAuth()).
        Use(middleware.CasbinHandler())
    {
        UserRouter.POST("changePassword", api.ChangePassword)     // 修密码
        UserRouter.POST("getUserList", api.GetUserList)           // 分页获取用户列表
        UserRouter.POST("setUserAuthority", api.SetUserAuthority) // 设置用户权限
        UserRouter.DELETE("deleteUser", api.DeleteUser)           // 删除用户
    }
}

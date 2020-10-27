package router

import (
   // "g-admin/middleware"
    "github.com/gin-gonic/gin"
   // "log"
    "net/http"
)

func InitUserRouter(Router *gin.RouterGroup) {
    UserRouter := Router.Group("user")
    {
        UserRouter.GET("/test", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "Hello world!",
            })
        })
        
        // UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
        // UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
        // UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
        // UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
        // UserRouter.PUT("setUserInfo", v1.SetUserInfo)            // 设置用户信息
    }
}
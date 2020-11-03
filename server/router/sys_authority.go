package router

import (
    "g-admin/middleware"
    "github.com/gin-gonic/gin"
    "g-admin/api"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
    AuthorityRouter := Router.Group("authority").
        Use(middleware.JWTAuth()).
        Use(middleware.CasbinHandler()).
        Use(middleware.GinDetailLogger())
    {
        AuthorityRouter.POST("createAuthority", api.CreateAuthority) // 创建角色
        AuthorityRouter.POST("deleteAuthority", api.DeleteAuthority) // 删除角色
        AuthorityRouter.PUT("updateAuthority", api.UpdateAuthority)  // 更新角色
        // AuthorityRouter.POST("copyAuthority", v1.CopyAuthority)       // 更新角色
        AuthorityRouter.POST("getAuthorityList", api.GetAuthorityList) // 获取角色列表
        // AuthorityRouter.POST("setDataAuthority", v1.SetDataAuthority) // 设置角色资源权限
    }
}

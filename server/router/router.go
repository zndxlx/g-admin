package router

import (
    "g-admin/core/log"
    "github.com/gin-gonic/gin"
    "g-admin/middleware"
)

func SetupRouter() *gin.Engine {

    r := gin.New()
    r.Use(middleware.GinRecovery(true))
    //Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
    // Router.Use(middleware.LoadTls())  // 打开就能玩https了
    //log.Info("use middleware logger")
    // 跨域
    //Router.Use(middleware.Cors())
    //log.Info("use middleware cors")

    // 方便统一添加路由组前缀 多服务器上线使用
    ApiGroup := r.Group("")
    InitBaseRouter(ApiGroup)
    InitUserRouter(ApiGroup)                  // 注册用户路由
    InitAuthorityRouter(ApiGroup)
    InitCustomerRouter(ApiGroup)
    InitMenuRouter(ApiGroup)
    log.Info("router register success")
    return r
}

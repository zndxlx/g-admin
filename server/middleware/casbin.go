package middleware


import (
    // "gin-vue-admin/global"
    // "gin-vue-admin/global/response"
    // "gin-vue-admin/model/request"
    // "gin-vue-admin/service"
    "fmt"
    "github.com/gin-gonic/gin"
    "g-admin/utils/response"
    "g-admin/service"
    "g-admin/utils/jwt"
    "g-admin/utils"
)

// 拦截器
func CasbinHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, _ := c.Get("claims")
        waitUse := claims.(*jwt.CustomClaims)
        // 获取请求的URI
        obj := c.Request.URL.RequestURI()
        // 获取请求方法
        act := c.Request.Method
        // 获取用户的角色
        sub := waitUse.AuthorityId
        e := service.Casbin()
        // 判断策略中是否存在
        success, err := e.Enforce(sub, obj, act)
        fmt.Printf("sub=%v, obj=%v, act=%v, success=%v \n", sub, obj, act, success)
        
        if success {
            c.Next()
        } else {
            utils.SetCtxExErr(c, err)
            response.Result(response.ERROR, gin.H{}, "权限不足,", c)
            c.Abort()
            return
        }
    }
}

package middleware

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "g-admin/utils/response"
    "time"
  //  "g-admin/service"
    "g-admin/utils/jwt"
    "g-admin/dao"
   // "g-admin/utils"
)

func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
        token := c.Request.Header.Get("x-token")
        if token == "" {
            response.Result(response.ERROR, gin.H{
                "reload": true,
            }, "未登录或非法访问", c)
            c.Abort()
            return
        }
        //j := NewJWT()
        // parseToken 解析token包含的信息
        claims, err := jwt.ParseToken(token)
        if err != nil {
            if err == jwt.TokenExpired {
                response.Result(response.ERROR, gin.H{
                    "reload": true,
                }, "授权已过期", c)
                c.Abort()
                return
            }
            response.Result(response.ERROR, gin.H{
                "reload": true,
            }, err.Error(), c)
            c.Abort()
            return
        }
        //判断用户是否被删除了
        if err, _ = dao.ISysUser.GetSysUserInfoByName(claims.Username); err != nil {
            response.Result(response.ERROR, gin.H{
                "reload": true,
            }, err.Error(), c)
            c.Abort()
        }
        
        //token自动续期
        if claims.ExpiresAt - time.Now().Unix()<claims.BufferTime {
            claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
            newToken,_ := jwt.CreateToken(*claims)
            newClaims,_ := jwt.ParseToken(newToken)
            c.Header("new-token",newToken)
            c.Header("new-expires-at",strconv.FormatInt(newClaims.ExpiresAt,10))
        }
        c.Set("claims", claims)
        c.Next()
    }
}

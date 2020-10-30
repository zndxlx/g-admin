package api

import (
    "g-admin/utils/response"
    "github.com/gin-gonic/gin"
    // "github.com/go-playground/validator/v10"
    // "g-admin/utils"
    "fmt"
    "g-admin/service"
    // "g-admin/dao"
    "github.com/dgrijalva/jwt-go"
    "time"
    "g-admin/utils/captcha"
    "g-admin/dao/model"
    jwt2 "g-admin/utils/jwt"
    "g-admin/dao"
    uuid "github.com/satori/go.uuid"
)

// User login structure
type LoginReq struct {
    Username  string `json:"username" binding:"required"`
    Password  string `json:"password" binding:"required"`
    Captcha   string `json:"captcha" binding:"required"`
    CaptchaId string `json:"captchaId" binding:"required"`
}

type LoginRsp struct {
    User      *model.SysUser `json:"user"`
    Token     string         `json:"token"`
    ExpiresAt int64          `json:"expiresAt"`
}

func Login(c *gin.Context) {
    var loginReq LoginReq
    err := c.ShouldBindJSON(&loginReq)
    if err != nil {
        errReqFormatRespone(err, c)
        return
    }
    captcha.Verify(loginReq.CaptchaId, loginReq.Captcha, true)
    if true {
        if err, user := service.Login(loginReq.Username, loginReq.Password); err != nil {
            response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
        } else {
            clams := jwt2.CustomClaims{
                UUID:        user.UUID,
                ID:          user.ID,
                NickName:    user.NickName,
                Username:    user.Username,
                AuthorityId: user.AuthorityId,
                BufferTime:  60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
                StandardClaims: jwt.StandardClaims{
                    NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
                    ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
                    Issuer:    "qmPlus",                       // 签名的发行者
                },
            }
            token, err := jwt2.CreateToken(clams)
            if err != nil {
                response.FailWithMessage("获取token失败", c)
                return
            }
            response.OkWithData(LoginRsp{
                User:      user,
                Token:     token,
                ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
            }, c)
            return
        }
        
    } else {
        response.FailWithMessage("验证码错误", c)
    }
}

type RegisterReq struct {
    HeaderImg   string `json:"headerImg" `
    NickName    string `json:"nickName" binding:"required"`
    Password    string `json:"password" binding:"required"`
    Username    string `json:"username" binding:"required"`
    AuthorityId string `json:"authorityId" binding:"required"`
}

type RegisterRsp struct {
    User model.SysUser `json:"user"`
}

func Register(c *gin.Context) {
    var R RegisterReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }
    
    user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
    err, userReturn := service.Register(user)
    if err != nil {
        response.FailWithDetailed(response.ERROR, RegisterRsp{User: *userReturn}, fmt.Sprintf("%v", err), c)
    } else {
        response.OkDetailed(RegisterRsp{User: *userReturn}, "注册成功", c)
    }
}


// Modify password structure
type ChangePasswordReq struct {
    Username    string `json:"username"`
    Password    string `json:"password"`
    NewPassword string `json:"newPassword"`
}

func ChangePassword(c *gin.Context) {
    var R ChangePasswordReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }

    U := &model.SysUser{Username: R.Username, Password: R.Password}
    if err, _ := service.ChangePassword(U, R.NewPassword); err != nil {
        response.FailWithMessage("修改失败，请检查用户名密码", c)
    } else {
        response.OkWithMessage("修改成功", c)
    }
}


func GetUserList(c *gin.Context) {
    var pageInfo model.PageInfo
    _ = c.ShouldBindJSON(&pageInfo)
    
    err, list, total := dao.ISysUser.GetUserInfoList(pageInfo)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
    } else {
        response.OkWithData(model.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, c)
    }
}

// Modify  user's auth structure
type SetUserAuthorityReq struct {
    UUID        uuid.UUID `json:"uuid"  binding:"required"`
    AuthorityId string    `json:"authorityId"  binding:"required"`
}

func SetUserAuthority(c *gin.Context) {
    var R SetUserAuthorityReq
    if err := c.ShouldBindJSON(&R); err != nil {
        errReqFormatRespone(err, c)
        return
    }
    
    err := dao.ISysUser.SetUserAuthority(R.UUID, R.AuthorityId)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("修改失败，%v", err), c)
    } else {
        response.OkWithMessage("修改成功", c)
    }
}


func DeleteUser(c *gin.Context) {
    var R model.GetById
    _ = c.ShouldBindJSON(&R)
    
    err := dao.ISysUser.DeleteUser(R.Id)
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
    } else {
        response.OkWithMessage("删除成功", c)
    }
}

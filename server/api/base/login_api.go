package base

import (
    "g-admin/utils/response"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "g-admin/utils"
    "fmt"
    "reflect"
)

// User login structure
type RegisterAndLoginStruct struct {
    Username  string `json:"username" binding:"required"`
    Password  string `json:"password" binding:"required"`
    Captcha   string `json:"captcha" binding:"required"`
    CaptchaId string `json:"captchaId" binding:"required"`
}

func Login(c *gin.Context) {
    var loginReq RegisterAndLoginStruct
    err := c.ShouldBindJSON(&loginReq)
    if err != nil {
        fmt.Println(reflect.TypeOf(err))
        errs, ok := err.(validator.ValidationErrors)
        if ok {
            response.FailWithData(utils.Translate(errs), c)
        }else{
            response.FailWithMessage(err.Error(), c)
        }
        
        return
    }

    if _store.Verify(loginReq.CaptchaId, loginReq.Captcha, true) {
        // U := &model.SysUser{Username: L.Username, Password: L.Password}
        // if err, user := service.Login(U); err != nil {
        //     response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), c)
        // } else {
        //     tokenNext(c, *user)
        // }
    } else {
        response.FailWithMessage("验证码错误", c)
    }
    
}
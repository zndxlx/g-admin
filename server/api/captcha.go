package api

import (
    "g-admin/utils/response"
    "fmt"
    "github.com/gin-gonic/gin"
    "g-admin/utils/captcha"
)

type SysCaptchaResponse struct {
    CaptchaId string `json:"captchaId"`
    PicPath   string `json:"picPath"`
}

func Captcha(c *gin.Context) {
    id, b64s, err := captcha.Generate()
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
    } else {
        response.OkDetailed(SysCaptchaResponse{
            CaptchaId: id,
            PicPath:   b64s,
        }, "验证码获取成功", c)
    }
}
package base

import (
    "g-admin/utils/response"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/mojocn/base64Captcha"
    "g-admin/config"
)

var _store = base64Captcha.DefaultMemStore
var _captchaConf *config.Captcha = &(config.Conf.Captcha)

type SysCaptchaResponse struct {
    CaptchaId string `json:"captchaId"`
    PicPath   string `json:"picPath"`
}

func Captcha(c *gin.Context) {
    //字符,公式,验证码配置
    // 生成默认数字的driver
    driver := base64Captcha.NewDriverDigit(_captchaConf.ImgHeight, _captchaConf.ImgWidth, _captchaConf.KeyLong, 0.7, 80)
    cp := base64Captcha.NewCaptcha(driver, _store)
    id, b64s, err := cp.Generate()
    if err != nil {
        response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
    } else {
        response.OkDetailed(SysCaptchaResponse{
            CaptchaId: id,
            PicPath:   b64s,
        }, "验证码获取成功", c)
    }
}
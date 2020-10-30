package captcha

import (
    "github.com/mojocn/base64Captcha"
    "g-admin/config"
)

var _store = base64Captcha.DefaultMemStore
var _capcha *base64Captcha.Captcha

func Init()  {
    var _captchaConf *config.Captcha = &(config.Conf.Captcha)
    driver := base64Captcha.NewDriverDigit(_captchaConf.ImgHeight, _captchaConf.ImgWidth, _captchaConf.KeyLong, 0.7, 80)
    _capcha = base64Captcha.NewCaptcha(driver, _store)
}

func Generate()(id, b64s string, err error) {
    return _capcha.Generate()
}

func Verify(id, answer string, clear bool) bool {
    return _store.Verify(id, answer, true)
}

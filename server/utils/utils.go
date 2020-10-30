package utils

import (
    "g-admin/utils/captcha"
    "g-admin/utils/jwt"
)

func Init() {
    initTrans("zh")
    captcha.Init()
    jwt.Init()
}
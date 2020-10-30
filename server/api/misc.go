package api

import (
    "github.com/go-playground/validator/v10"
    "g-admin/utils/response"
    "g-admin/utils"
    "github.com/gin-gonic/gin"
)



func errReqFormatRespone(err error, c *gin.Context) {
    errs, ok := err.(validator.ValidationErrors)
    if ok {
        response.FailWithData(utils.Translate(errs), c)
    }else{
        response.FailWithMessage(err.Error(), c)
    }
}

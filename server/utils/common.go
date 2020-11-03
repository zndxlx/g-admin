package utils

import (
    "github.com/gin-gonic/gin"
    "fmt"
    //"github.com/pkg/errors"
)

const ExtLogKey = "extlog"
func SetCtxExErr(c *gin.Context, err error){
    c.Set(ExtLogKey, fmt.Sprintf("%+v", err))
}

func GetCtxExErr(c *gin.Context, ) string{
    if s, ok := c.Get(ExtLogKey); ok {
        return s.(string)
    }else{
        return ""
    }
}
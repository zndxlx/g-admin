package main

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    
    "g-admin/config"
    "time"
    "g-admin/core"
    "g-admin/router"
)

func initServer(address string, router *gin.Engine) *http.Server {
    return &http.Server{
        Addr:           address,
        Handler:        router,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
}


func main() {
    config.InitConfig("./config.yaml")
    core.InitCore()

    eg := router.SetupRouter()
    server := initServer(config.Conf.Base.Addr,eg)
    server.ListenAndServe()
}

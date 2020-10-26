package main

import (
    "g-admin/config"
    "g-admin/core"
    "g-admin/core/log"
    "go.uber.org/zap"
)

func main() {
    
    config.InitConfig("./config.yaml")
    core.InitCore()
    log.Info("Success..",
        zap.String("statusCode", "444"),
        zap.String("url", "test_url"))
}
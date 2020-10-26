package config

import (
    "github.com/spf13/viper"
    "fmt"
    "github.com/fsnotify/fsnotify"
)


type Config struct {
    Redis Redis `mapstructure:"redis"`
    Mysql Mysql `mapstructure:"mysql"`
    Zap  Zap `mapstructure:"zap"`
}

type Zap struct {
    Level     string `mapstructure:"addr"`
    Path     string `mapstructure:"path"`
    LogInConsole     bool `mapstructure:"log-in-console"`
    MaxSize   int  `mapstructure:"max-size"`
    MaxBackups   int  `mapstructure:"max-backups"`
}

type Redis struct {
    Addr     string `mapstructure:"addr"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type Mysql struct {
    Username     string `mapstructure:"username" `
    Password     string `mapstructure:"password"`
    Dbname       string `mapstructure:"db-name"`
    Path         string `mapstructure:"path"`
    Config       string `mapstructure:"config"`
    MaxIdleConns int    `mapstructure:"max-idle-conns"`
    MaxOpenConns int    `mapstructure:"max-open-conns"`
}

var Conf = &Config{
    Redis: Redis{
        DB: 10, // 设置默认值
    },
}

func InitConfig(path string) {
    viper.SetConfigFile(path)   // 指定配置文件路径
    err := viper.ReadInConfig() // 读取配置信息
    if err != nil { // 读取配置信息失败
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    // 将读取的配置信息保存至全局变量Conf
    if err := viper.Unmarshal(Conf); err != nil {
        panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
    }
    // 监控配置文件变化
    viper.WatchConfig()
    // 注意！！！配置文件发生变化后要同步到全局变量Conf
    viper.OnConfigChange(func(in fsnotify.Event) {
        fmt.Println("配置文件被人修改啦...")
        if err := viper.Unmarshal(Conf); err != nil {
            panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
        }
        fmt.Printf("%+v\n", *Conf)
    })
    fmt.Printf("%+v\n", *Conf)
}

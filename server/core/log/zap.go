package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/natefinch/lumberjack.v2"
    "g-admin/config"
    "os"
)

var _logger *zap.Logger
var _zapConf  config.Zap

func getlogLevel(s string)( level zapcore.Level){
    switch s { // 初始化配置文件的Level
    case "debug":
        level = zap.DebugLevel
    case "info":
        level = zap.InfoLevel
    case "warn":
        level = zap.WarnLevel
    case "error":
        level = zap.ErrorLevel
    case "dpanic":
        level = zap.DPanicLevel
    case "panic":
        level = zap.PanicLevel
    case "fatal":
        level = zap.FatalLevel
    default:
        level = zap.InfoLevel
    }
    return
}

func InitLogger() {
    _zapConf = config.Conf.Zap
    writeSyncer := getLogWriter()
    encoder := getEncoder()
    level := getlogLevel(_zapConf.Level)
    core := zapcore.NewCore(encoder, writeSyncer, level)
    _logger = zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1))

}

func getEncoder() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    //encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
    lumberJackLogger := &lumberjack.Logger{
        Filename:   _zapConf.Path,
        MaxSize:    _zapConf.MaxSize,   // 单位MB
        MaxBackups: _zapConf.MaxBackups,
        MaxAge:     30,
        Compress:   false,
    }
    if _zapConf.LogInConsole {
        return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
    }
    
    return zapcore.AddSync(lumberJackLogger)
}


func Info(msg string, fields ...zap.Field) {
    _logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
    _logger.Warn(msg, fields...)
}


func Error(msg string, fields ...zap.Field) {
    _logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
    _logger.Panic(msg, fields...)
}


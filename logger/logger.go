package logger

import (
	"fmt"
	"github.com/huyujie/gin-demo/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func Setup() {

	c := config.ReadConf()

	lumlog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   c.Log.Filename,
		MaxSize:    c.Log.MaxSize,
		MaxBackups: c.Log.MaxBackups,
		MaxAge:     c.Log.MaxAge,
	})

	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.TimeKey = "time"
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(EncoderConfig),
		lumlog,
		zap.InfoLevel,
	)
	Logger = zap.New(core)
	defer Logger.Sync()
}

func Debug(msg string, fields zap.Field) {
	Logger.Debug(msg, fields)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Debugf(format string, v ...interface{}) {
	Logger.Debug(fmt.Sprintf(format, v))
}

func Infof(format string, v ...interface{}) {
	Logger.Info(fmt.Sprintf(format, v))
}

func Warnf(format string, v ...interface{}) {
	Logger.Warn(fmt.Sprintf(format, v))
}

func Errorf(format string, v ...interface{}) {
	Logger.Error(fmt.Sprintf(format, v))
}

func Fatalf(format string, v ...interface{}) {
	Logger.Fatal(fmt.Sprintf(format, v))
}

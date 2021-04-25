// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package logger

import (
	"fmt"
	"gin-demo/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

var logger *zap.Logger

func GetLogger() *zap.Logger {
	return logger
}

func Setup(ServiceName string) {

	config := config.ReadConf()

	//convert gin-demo to gin_demo
	LogName := strings.Replace(ServiceName, "-", "_", 1)

	FileName := fmt.Sprintf("%s/%s.log", config.Common.Log.Filepath, LogName)

	lumlog := zapcore.AddSync(&lumberjack.Logger{
		Filename:   FileName,
		MaxSize:    config.Common.Log.MaxSize,
		MaxBackups: config.Common.Log.MaxBackups,
		MaxAge:     config.Common.Log.MaxAge,
	})

	EncoderConfig := zap.NewProductionEncoderConfig()
	EncoderConfig.TimeKey = "time"
	EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	level := GetLogLevel(config.Common.Log.Level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(EncoderConfig),
		lumlog,
		level,
	)
	logger = zap.New(core, zap.AddCaller()).WithOptions(zap.AddCallerSkip(1))
	defer logger.Sync()
}

func GetLogLevel(level string) zapcore.Level {
	if level == "info" {
		return zap.InfoLevel
	} else if level == "debug" {
		return zap.DebugLevel
	} else if level == "error" {
		return zap.ErrorLevel
	}
	return zap.InfoLevel
}

func Debug(msg string, fields zap.Field) {
	logger.Debug(msg, fields)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	logger.Info(fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	logger.Warn(fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, v...))
}

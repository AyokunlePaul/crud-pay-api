package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
)

var zapLogger *zap.Logger
var zapLoggerConfig zap.Config

func init() {
	setup()
}

func setup() {
	if zapLogger == nil {
		zapLoggerConfig = zap.Config{
			OutputPaths: []string{"stdout"},
			Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
			Encoding:    "json",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey:     "message",
				LevelKey:       "level",
				TimeKey:        "time",
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
			},
		}
		var zapInitError error
		if zapLogger, zapInitError = zapLoggerConfig.Build(); zapInitError != nil {
			panic(zapInitError)
		}
	}
}

func GetLogger() *zap.Logger {
	if zapLogger == nil {
		setup()
	}
	return zapLogger
}

func Info(message string, tags ...zap.Field) {
	zapLogger.Info(message, tags...)
	_ = zapLogger.Sync()
}

func Error(message string, logError error, tags ...zap.Field) {
	tags = append(tags, zap.String("name", reflect.TypeOf(logError).Name()), zap.NamedError("error", logError))
	zapLogger.Error(message, tags...)
	_ = zapLogger.Sync()
}

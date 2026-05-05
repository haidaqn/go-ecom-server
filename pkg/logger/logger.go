package logger

import (
	"os"

	"github.com/haidaqn/go-ecommerce-backend-api/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Level
	var level zapcore.Level

	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderConfig()
	hook := lumberjack.Logger{
		Filename:   config.File,
		MaxSize:    config.MaxSize,
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,
		Compress:   config.Compress,
	}
	writer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook), zapcore.AddSync(os.Stdout))
	core := zapcore.NewCore(encoder, writer, level)
	logger := &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))}
	return logger
}

func getEncoderConfig() zapcore.Encoder {
	endcodeConfig := zap.NewProductionEncoderConfig()
	endcodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	endcodeConfig.TimeKey = "time"

	endcodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	endcodeConfig.EncodeDuration = zapcore.StringDurationEncoder
	endcodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	endcodeConfig.EncodeName = zapcore.FullNameEncoder

	return zapcore.NewJSONEncoder(endcodeConfig)
}

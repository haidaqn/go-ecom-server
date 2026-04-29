package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderConfig()
	writer := getWriterSync()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info level log", zap.String("key", "value"))
	logger.Debug("Debug level log", zap.String("key", "value"))
	logger.Error("Error level log", zap.String("key", "value"))
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

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	syncCore := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncCore, syncConsole)
}

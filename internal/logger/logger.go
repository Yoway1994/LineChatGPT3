package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	now := time.Now()
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("log/%04d-%02d-%02d.log", now.Year(), now.Month(), now.Day()),
		MaxSize:    500,
		MaxBackups: 2,
		MaxAge:     30,
	})
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.InfoLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}

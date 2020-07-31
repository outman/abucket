package pkg

import (
	"sync"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	zapLogger *zap.Logger
	onceZap   sync.Once
)

// NewZapLogger *zap.Logger
func NewZapLogger() *zap.Logger {
	onceZap.Do(func() {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   viper.GetString("LOG_PATH"),
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     7,
		})
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		)
		zapLogger = zap.New(core)
	})

	return zapLogger
}

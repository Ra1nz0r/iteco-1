package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapService interface {
	Info(fields ...interface{})
	Error(fields ...interface{})
	Fatal(fields ...interface{})
}

type ZapStorage struct {
	*zap.Logger
}

var Zap ZapService = &ZapStorage{zap.NewNop()}

func Initialize() error {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	config.DisableStacktrace = true

	logger, err := config.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return fmt.Errorf("logger build error: %w", err)
	}

	Zap = &ZapStorage{logger}

	return nil
}

func (z *ZapStorage) Info(fields ...interface{}) {
	z.Logger.Sugar().Infoln(fields...)
}

func (z *ZapStorage) Error(fields ...interface{}) {
	z.Logger.Sugar().Errorln(fields...)
}

func (z *ZapStorage) Fatal(fields ...interface{}) {
	z.Logger.Sugar().Fatalln(fields...)
}

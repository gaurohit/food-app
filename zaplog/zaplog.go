package zaplog

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLog(levels zapcore.Level) (*zap.Logger, error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(levels),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}
	logger, err := config.Build()
	if err != nil {
		log.Println(err)
	}
	return logger, nil
}

func ZapcoreFunc(key string, body string, Interface interface{}) zapcore.Field {
	bodyField := zapcore.Field{
		Key:       key,
		Type:      zapcore.StringType,
		String:    body,
		Interface: nil,
	}
	return bodyField
}

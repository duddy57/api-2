package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(env string) (l *zap.Logger, err error) {
	switch strings.ToLower(env) {
	case "production":
		l, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	case "development":
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		l, err = cfg.Build()
		if err != nil {
			return nil, err
		}
	default:
		l, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}

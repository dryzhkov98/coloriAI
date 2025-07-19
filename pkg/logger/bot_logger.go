package logger

import "go.uber.org/zap"

type ZapLoggerAdapter struct {
	sugar *zap.SugaredLogger
}

func NewAdapter(logger *zap.Logger) *ZapLoggerAdapter {
	return &ZapLoggerAdapter{
		sugar: logger.Sugar(),
	}
}

func (z *ZapLoggerAdapter) Println(v ...interface{}) {
	z.sugar.Info(v...)
}

func (z *ZapLoggerAdapter) Printf(format string, v ...interface{}) {
	z.sugar.Infof(format, v...)
}

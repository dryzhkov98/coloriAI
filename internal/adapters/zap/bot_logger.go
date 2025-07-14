package zapbotlogger

import "coloriAI/pkg/logger"

type ZapLoggerAdapter struct{}

func (z *ZapLoggerAdapter) Println(v ...interface{}) {
	logger.Get().Sugar().Info(v...)
}

func (z *ZapLoggerAdapter) Printf(format string, v ...interface{}) {
	logger.Get().Sugar().Infof(format, v...)
}

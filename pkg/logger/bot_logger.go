package logger

// перенести в pkg
type ZapLoggerAdapter struct{}

func (z *ZapLoggerAdapter) Println(v ...interface{}) {
	Get().Sugar().Info(v...)
}

func (z *ZapLoggerAdapter) Printf(format string, v ...interface{}) {
	Get().Sugar().Infof(format, v...)
}

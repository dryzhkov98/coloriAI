package logger

import (
	"coloriAI/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New(cfg *config.Config) *zap.Logger {
	logger := newLogger(cfg)
	zap.ReplaceGlobals(logger)

	return logger
}

func newLogger(config *config.Config) *zap.Logger {
	loglevel := parseLogLevel(config.AppConfig.LogLevel)

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		MessageKey:  "msg",
		EncodeTime:  zapcore.ISO8601TimeEncoder,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	stdoutCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zapcore.ErrorLevel && l >= loglevel
	}))

	stderrCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stderr), zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zapcore.ErrorLevel && l >= loglevel
	}))

	core := zapcore.NewTee(stdoutCore, stderrCore)

	additionalFields := zap.Fields(
		zap.String("application", config.BotConfig.BotName),
		zap.String("environment", config.AppConfig.Environment),
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		additionalFields,
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return logger
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

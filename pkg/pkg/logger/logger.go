package logger

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func InitLogging() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	Logger = logger.Sugar()
}

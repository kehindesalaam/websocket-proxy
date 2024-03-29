package log

import "go.uber.org/zap"

func Error(args ...interface{}) {
	zap.S().Error(args)
}

func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args)
}

func Info(args ...interface{}) {
	zap.S().Info(args)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args)
}

func Warn(args ...interface{}) {
	zap.S().Warn(args)
}

func Warnf(template string, args ...interface{}) {
	zap.S().Warnf(template, args)
}

func Fatal(args ...interface{}) {
	zap.S().Fatal(args)
}

func Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args)
}
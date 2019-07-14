package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

type logger struct {
	proxy *zap.SugaredLogger
}

func New(fileName string, level string, maxSize int, maxBackups int, maxAge int) *logger {
	log := &logger{}
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  true,
		Compress:   true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(getLoggerLevel(level)))
	zap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log.proxy = zap.Sugar()
	return log
}

func (log *logger) Debug(args ...interface{}) {
	log.proxy.Debug(args...)
}

func (log *logger) Debugf(template string, args ...interface{}) {
	log.proxy.Debugf(template, args...)
}

func (log *logger) Info(args ...interface{}) {
	log.proxy.Info(args...)
}

func (log *logger) Infof(template string, args ...interface{}) {
	log.proxy.Infof(template, args...)
}

func (log *logger) Warn(args ...interface{}) {
	log.proxy.Warn(args...)
}

func (log *logger) Warnf(template string, args ...interface{}) {
	log.proxy.Warnf(template, args...)
}

func (log *logger) Error(args ...interface{}) {
	log.proxy.Error(args...)
}

func (log *logger) Errorf(template string, args ...interface{}) {
	log.proxy.Errorf(template, args...)
}

func (log *logger) DPanic(args ...interface{}) {
	log.proxy.DPanic(args...)
}

func (log *logger) DPanicf(template string, args ...interface{}) {
	log.proxy.DPanicf(template, args...)
}

func (log *logger) Panic(args ...interface{}) {
	log.proxy.Panic(args...)
}

func (log *logger) Panicf(template string, args ...interface{}) {
	log.proxy.Panicf(template, args...)
}

func (log *logger) Fatal(args ...interface{}) {
	log.proxy.Fatal(args...)
}

func (log *logger) Fatalf(template string, args ...interface{}) {
	log.proxy.Fatalf(template, args...)
}

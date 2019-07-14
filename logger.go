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

type Logger struct {
	proxy *zap.SugaredLogger
}

func New(fileName string, level string, maxSize int, maxBackups int, maxAge int) *Logger {
	log := &Logger{}
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

func (log *Logger) Debug(args ...interface{}) {
	log.proxy.Debug(args...)
}

func (log *Logger) Debugf(template string, args ...interface{}) {
	log.proxy.Debugf(template, args...)
}

func (log *Logger) Info(args ...interface{}) {
	log.proxy.Info(args...)
}

func (log *Logger) Infof(template string, args ...interface{}) {
	log.proxy.Infof(template, args...)
}

func (log *Logger) Warn(args ...interface{}) {
	log.proxy.Warn(args...)
}

func (log *Logger) Warnf(template string, args ...interface{}) {
	log.proxy.Warnf(template, args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.proxy.Error(args...)
}

func (log *Logger) Errorf(template string, args ...interface{}) {
	log.proxy.Errorf(template, args...)
}

func (log *Logger) DPanic(args ...interface{}) {
	log.proxy.DPanic(args...)
}

func (log *Logger) DPanicf(template string, args ...interface{}) {
	log.proxy.DPanicf(template, args...)
}

func (log *Logger) Panic(args ...interface{}) {
	log.proxy.Panic(args...)
}

func (log *Logger) Panicf(template string, args ...interface{}) {
	log.proxy.Panicf(template, args...)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.proxy.Fatal(args...)
}

func (log *Logger) Fatalf(template string, args ...interface{}) {
	log.proxy.Fatalf(template, args...)
}

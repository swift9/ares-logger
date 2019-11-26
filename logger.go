package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"DEBUG":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"INFO":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"WARN":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"ERROR":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"DPANIC": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"PANIC":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
	"FATAL":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

type ILogger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(template string, args ...interface{})
}

type Logger struct {
	ZapSugared *zap.SugaredLogger
}

func New(fileName string, level string, maxSize int, maxBackups int, maxAge int) *Logger {
	zapLog := &Logger{}
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
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(getLoggerLevel(level)))
	zap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog.ZapSugared = zap.Sugar()
	return zapLog
}

func (log *Logger) Debug(args ...interface{}) {
	log.ZapSugared.Debug(args...)
}

func (log *Logger) Debugf(template string, args ...interface{}) {
	log.ZapSugared.Debugf(template, args...)
}

func (log *Logger) Debugw(template string, args ...interface{}) {
	log.ZapSugared.Debugw(template, args...)
}

func (log *Logger) Info(args ...interface{}) {
	log.ZapSugared.Info(args...)
}

func (log *Logger) Infof(template string, args ...interface{}) {
	log.ZapSugared.Infof(template, args...)
}

func (log *Logger) Infow(template string, args ...interface{}) {
	log.ZapSugared.Infow(template, args...)
}

func (log *Logger) Warn(args ...interface{}) {
	log.ZapSugared.Warn(args...)
}

func (log *Logger) Warnf(template string, args ...interface{}) {
	log.ZapSugared.Warnf(template, args...)
}

func (log *Logger) Warnw(template string, args ...interface{}) {
	log.ZapSugared.Warnw(template, args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.ZapSugared.Error(args...)
}

func (log *Logger) Errorf(template string, args ...interface{}) {
	log.ZapSugared.Errorf(template, args...)
}

func (log *Logger) Errorw(template string, args ...interface{}) {
	log.ZapSugared.Errorw(template, args...)
}

type JsonLogger struct {
	ZapSugared *zap.SugaredLogger
}

func NewJson(fileName string, level string, maxSize int, maxBackups int, maxAge int) *JsonLogger {
	zapLog := &JsonLogger{}
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
	zapLog.ZapSugared = zap.Sugar()
	return zapLog
}

func NewJsonWithMsgKey(fileName string, level string, maxSize int, maxBackups int, maxAge int, msgKey string) *JsonLogger {
	zapLog := &JsonLogger{}
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  true,
		Compress:   true,
	})
	encoder := zap.NewProductionEncoderConfig()
	if msgKey != "" {
		encoder.MessageKey = msgKey
	}
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(getLoggerLevel(level)))
	zap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog.ZapSugared = zap.Sugar()
	return zapLog
}

func (log *JsonLogger) Debug(args ...interface{}) {
	log.ZapSugared.Debug(args...)
}

func (log *JsonLogger) Debugf(template string, args ...interface{}) {
	log.ZapSugared.Debugf(template, args...)
}

func (log *JsonLogger) Debugw(template string, args ...interface{}) {
	log.ZapSugared.Debugw(template, args...)
}

func (log *JsonLogger) Info(args ...interface{}) {
	log.ZapSugared.Info(args...)
}

func (log *JsonLogger) Infof(template string, args ...interface{}) {
	log.ZapSugared.Infof(template, args...)
}

func (log *JsonLogger) Infow(template string, args ...interface{}) {
	log.ZapSugared.Infow(template, args...)
}

func (log *JsonLogger) Warn(args ...interface{}) {
	log.ZapSugared.Warn(args...)
}

func (log *JsonLogger) Warnf(template string, args ...interface{}) {
	log.ZapSugared.Warnf(template, args...)
}

func (log *JsonLogger) Warnw(template string, args ...interface{}) {
	log.ZapSugared.Warnw(template, args...)
}

func (log *JsonLogger) Error(args ...interface{}) {
	log.ZapSugared.Error(args...)
}

func (log *JsonLogger) Errorf(template string, args ...interface{}) {
	log.ZapSugared.Errorf(template, args...)
}

func (log *JsonLogger) Errorw(template string, args ...interface{}) {
	log.ZapSugared.Errorw(template, args...)
}

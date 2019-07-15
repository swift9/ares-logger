package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
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
	ZapSugared *zap.SugaredLogger
}

func New(fileName string, level string, maxSize int, maxBackups int, maxAge int) *Logger {
	log.Println("init zap log ...")
	log.Println("fileName: "+fileName, " level:", level, " maxSize:", maxSize, " maxBackups:", maxBackups, " maxAge:", maxAge)
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
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(getLoggerLevel(level)))
	zap := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog.ZapSugared = zap.Sugar()
	log.Println("zap log is ready")
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

func (log *Logger) DPanic(args ...interface{}) {
	log.ZapSugared.DPanic(args...)
}

func (log *Logger) DPanicf(template string, args ...interface{}) {
	log.ZapSugared.DPanicf(template, args...)
}

func (log *Logger) DPanicw(template string, args ...interface{}) {
	log.ZapSugared.DPanicw(template, args...)
}

func (log *Logger) Panic(args ...interface{}) {
	log.ZapSugared.Panic(args...)
}

func (log *Logger) Panicf(template string, args ...interface{}) {
	log.ZapSugared.Panicf(template, args...)
}

func (log *Logger) Panicw(template string, args ...interface{}) {
	log.ZapSugared.Panicw(template, args...)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.ZapSugared.Fatal(args...)
}

func (log *Logger) Fatalf(template string, args ...interface{}) {
	log.ZapSugared.Fatalf(template, args...)
}

func (log *Logger) Fatalw(template string, args ...interface{}) {
	log.ZapSugared.Fatalw(template, args...)
}

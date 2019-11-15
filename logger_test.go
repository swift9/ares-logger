package logger_test

import (
	"errors"
	logger "github.com/swift9/ares-logger"
	"log"
	"testing"
)

var Log *logger.Logger

var JsonLog *logger.JsonLogger

func init() {
	log.Println("init logger Log")
	Log = logger.New(
		"app.log",
		"INFO",
		1024,
		30,
		30)

	JsonLog = logger.NewJson(
		"app.json.log",
		"INFO",
		1024,
		30,
		30)
}

func TestLog(t *testing.T) {
	Log.Info("log is ready", "a", 1, "b", 2, "error", errors.New("ss"))
}

func TestLogw(t *testing.T) {
	Log.Infow("log is ready", "a", 1, "b", 2, "error", errors.New("ss"))
}

func TestJsonLog(t *testing.T) {
	JsonLog.Info("log is ready", "a", 1, "b", 2, "error", errors.New("ss"))
}

func TestJsonLogw(t *testing.T) {
	JsonLog.Infow("log is ready", "a", 1, "b", 2, "error", errors.New("ss"))
}

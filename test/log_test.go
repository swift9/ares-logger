package test

import (
	logger "github.com/swift9/ares-logger"
	"testing"
)

var log = logger.New("test.log", "info", 1024, 20, 30)

func TestInfo(t *testing.T) {
	log.Info("test")
}

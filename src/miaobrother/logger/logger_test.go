package logger

import (
	"testing"
)

func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("User is come from china,the id is %d\n", 12345)
	logger.Warn("test log in warn\n")
	logger.Fatal("This is a fatal\n")
	logger.Close()

}

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "/Users/zhangyalei", "test")
	logger.Debug("User is come from china,the id is %d\n", 12345)
	logger.Warn("test log in warn\n")
	logger.Fatal("This is a fatal\n")
	logger.Close()
}

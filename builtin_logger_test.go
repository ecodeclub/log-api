package logs

import (
	"log"
	"testing"
)

// 单纯提高一些覆盖率
func TestBuiltInLogger(t *testing.T) {
	logger := NewBuiltInLogger(log.Default())
	t.Run("debug", func(t *testing.T) {
		logger.Debug("hello, %s", "world")
	})
	t.Run("info", func(t *testing.T) {
		logger.Info("hello, %s", "world")
	})
	t.Run("warn", func(t *testing.T) {
		logger.Warn("hello, %s", "world")
	})
	t.Run("error", func(t *testing.T) {
		logger.Error("hello, %s", "world")
	})
}

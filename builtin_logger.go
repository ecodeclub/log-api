package logs

import "log"

type BuiltInLogger struct {
	logger *log.Logger
}

func NewBuiltInLogger(l *log.Logger) *BuiltInLogger {
	return &BuiltInLogger{logger: l}
}

func (b *BuiltInLogger) Debug(msg string, args ...any) {
	b.logger.Printf(msg, args...)
}

func (b *BuiltInLogger) Info(msg string, args ...any) {
	b.logger.Printf(msg, args...)
}

func (b *BuiltInLogger) Warn(msg string, args ...any) {
	b.logger.Printf(msg, args...)
}

func (b *BuiltInLogger) Error(msg string, args ...any) {
	b.logger.Printf(msg, args...)
}

func (b *BuiltInLogger) Fatal(msg string, args ...any) {
	b.logger.Fatalf(msg, args...)
}

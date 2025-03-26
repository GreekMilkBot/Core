package Core

import (
	"context"
	"log"
)

type LoggerFactory interface {
	Logger(ctx context.Context) Logger
}

type Logger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

type noopLoggerFactory struct {
	Group string
}

func (*noopLoggerFactory) Logger(ctx context.Context) Logger {
	if group, ok := ctx.Value("module-name").(string); ok {
		return noopLogger{group}
	}
	return noopLogger{"root"}
}

type noopLogger struct {
	group string
}

func (n noopLogger) Debugf(format string, args ...any) {
	log.Printf(" [DEBUG] ["+n.group+"] "+format, args...)
}

func (n noopLogger) Infof(format string, args ...any) {
	log.Printf(" [INFO ] ["+n.group+"] "+format, args...)
}

func (n noopLogger) Warnf(format string, args ...any) {
	log.Printf(" [WARN ] ["+n.group+"] "+format, args...)
}

func (n noopLogger) Errorf(format string, args ...any) {
	log.Printf(" [ERROR ] ["+n.group+"] "+format, args...)
}

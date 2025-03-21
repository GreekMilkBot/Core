package Core

import (
	"context"
)

type Context struct {
	context.Context
	Logger
}

func NewContext(ctx context.Context) Context {
	logger := ctx.Value("logger").(LoggerFactory)
	if logger == nil {
		logger = &noopLoggerFactory{}
		ctx = context.WithValue(ctx, "logger", logger)
	}

	return Context{
		Context: ctx,
		Logger:  logger.Logger(ctx),
	}
}

func (ctx Context) WithValue(key, value any) Context {
	return Context{
		Context: context.WithValue(ctx.Context, key, value),
		Logger:  ctx.Value("logger").(LoggerFactory).Logger(ctx.Context),
	}
}

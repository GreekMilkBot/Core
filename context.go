package Core

import (
	"context"

	"code.d7z.net/d7z-team/go-variables"
)

type Context struct {
	context.Context
	Logger
	Config variables.Variables
}

type BotInstances struct {
	plugins map[string][]*Module
}

func (ctx Context) Modules(mod string) []*Module {
	if find, ok := ctx.Value(mod).([]*Module); ok {
		return find
	}
	return []*Module{}
}

func NewContext(ctx context.Context) Context {
	if _, ok := ctx.Value("logger").(LoggerFactory); !ok {
		ctx = context.WithValue(ctx, "logger", &noopLoggerFactory{})
	}

	return Context{
		Context: ctx,
		Logger:  ctx.Value("logger").(LoggerFactory).Logger(ctx),
		Config:  variables.Variables{},
	}
}

func (ctx Context) WithValue(key, value any) Context {
	nextCtx := context.WithValue(ctx.Context, key, value)
	return Context{
		Context: ctx,
		Logger:  ctx.Value("logger").(LoggerFactory).Logger(nextCtx),
	}
}

package share

import "context"

type Context struct {
	context.Context
	Logger
}

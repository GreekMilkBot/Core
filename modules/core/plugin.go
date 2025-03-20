package core

import (
	"github.com/GreekMilkBot/Core/modules/core/share"
)

type Plugin interface {
	Name() string
}

type PluginActions interface {
	OnStart(ctx share.Context) error
	OnDestroy(ctx share.Context) error
}

type AdapterActions interface{}

type Adapter interface {
	Plugin
	LoopReceive(ctx share.Context) error
	Actions() AdapterActions
}

type Client interface {
	Plugin
	OnMessage(ctx share.Context, args interface{}) error
}

package core

import (
	"github.com/GreekMilkBot/Core/modules/core/share"
)

type Plugin interface {
	Name() string
}

type PluginActions interface {
	Start() error
	Stop() error
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

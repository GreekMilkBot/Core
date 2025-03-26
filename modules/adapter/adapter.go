package adapter

import (
	"github.com/GreekMilkBot/Core"
)

func init() {
	Core.RegisterModule(Adapter{})
}

type Adapter struct{}

func (a Adapter) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID: "system.adapter",
		New: func() Core.Module {
			return new(Adapter)
		},
	}
}

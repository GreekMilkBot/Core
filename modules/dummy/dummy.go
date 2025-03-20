package dummy

import "github.com/GreekMilkBot/Core"

func init() {
	Core.RegisterModule(Dummy{})
}

type Dummy struct{}

func (d Dummy) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID: "adapter.dummy",
		New: func() Core.Module {
			return Dummy{}
		},
	}
}

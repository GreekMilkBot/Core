package onebot

import "github.com/GreekMilkBot/Core"

func init() {
	Core.RegisterModule(OneBot{})
}

type OneBot struct{}

func (o OneBot) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID:  "adapter.onebot",
		New: func() Core.Module { return new(OneBot) },
	}
}

package logging

import (
	"github.com/GreekMilkBot/Core"
)

func init() {
	Core.RegisterModule(ZapLogger{})
}

type ZapLogger struct {
}

func (z ZapLogger) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID: "system.logging",
		New: func() Core.Module {
			return new(ZapLogger)
		},
	}
}

package logging

import (
	"github.com/GreekMilkBot/Core"
)

func init() {
	Core.RegisterModule(Logger{})
}

type Logger struct{}

func (z Logger) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID: "system.logging",
		New: func() Core.Module {
			return new(Logger)
		},
	}
}

func (z Logger) Priority() int {
	return 200
}

func (z Logger) Processor(context *Core.Context, mods *[]*Core.BotInstance) error {
	return nil
}

func (z Logger) Provision(ctx Core.Context) error {
	ctx.Infof("[Logger] Init Logger")
	return nil
}

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

func (z ZapLogger) Priority() int {
	return 200
}
func (z ZapLogger) Processor(context *Core.Context, mods *[]*Core.BotInstance) error {
	return nil
}

func (z ZapLogger) Provision(ctx Core.Context) error {
	ctx.Infof("[ZapLogger] Init ZapLogger")
	return nil
}

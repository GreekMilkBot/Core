package order

import (
	"github.com/GreekMilkBot/Core"
)

type Order struct {
}

func (o *Order) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID: "system.order",
		New: func() Core.Module {
			return new(Order)
		},
	}
}

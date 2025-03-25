package order

import (
	"github.com/GreekMilkBot/Core"
	"slices"
	"strings"
)

func init() {
	Core.RegisterModule(Order{})
}

type Order struct {
}

func nameId(key string) (int, string) {
	if before, after, found := strings.Cut(key, "."); found {
		switch before {
		case "system":
			return 0, after
		case "admin":
			return 1, after
		case "adapter":
			return 2, after
		}
	}
	return 10, key
}

func (o Order) Processor(_ *Core.Context, mods *[]Core.BotInstance) error {
	slices.SortFunc(*mods, func(a, b Core.BotInstance) int {
		idA, bodyA := nameId(a.ID)
		idB, bodyB := nameId(b.ID)
		if idA != idB {
			return idA - idB
		}
		return strings.Compare(bodyA, bodyB)
	})
	return nil
}

func (o Order) Provision(ctx Core.Context) error {
	ctx.Infof("Order.Provision")
	return nil
}

func (o Order) Priority() int {
	return 100
}

func (o Order) BotModule() Core.ModuleInfo {
	return Core.ModuleInfo{
		ID: "system.order",
		New: func() Core.Module {
			return new(Order)
		},
	}
}

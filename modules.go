package Core

import "fmt"

func RegisterModule(instance Module) {
	module := instance.BotModule()
	fmt.Printf("RegisterModule %s\n", module.ID)
	//todo
}

type Module interface {
	BotModule() ModuleInfo
}

type ModuleInfo struct {
	ID  string
	New func() Module
}

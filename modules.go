package Core

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type Module interface {
	BotModule() ModuleInfo
}

type ModuleInfo struct {
	ID    string // 插件 ID
	Multi bool   // 是否支持多实例
	New   func() Module
}

func RegisterModule(instance Module) {
	mod := instance.BotModule()

	if mod.ID == "" {
		panic("module ID missing")
	}
	if mod.New == nil {
		panic("missing ModuleInfo.New")
	}
	if val := mod.New(); val == nil {
		panic("ModuleInfo.New must return a non-nil module instance")
	}

	modulesMu.Lock()
	defer modulesMu.Unlock()

	if _, ok := modules[mod.ID]; ok {
		panic(fmt.Sprintf("module already registered: %s", mod.ID))
	}
	modules[mod.ID] = mod
}

func Modules(prefix string) []string {
	modulesMu.RLock()
	defer modulesMu.RUnlock()

	names := make([]string, 0, len(modules))
	for name := range modules {
		if strings.HasPrefix(name, prefix) {
			names = append(names, name)
		}
	}

	sort.Strings(names)

	return names
}

// GetModule returns module information from its ID (full name).
func GetModule(name string) (ModuleInfo, error) {
	modulesMu.RLock()
	defer modulesMu.RUnlock()
	m, ok := modules[name]
	if !ok {
		return ModuleInfo{}, fmt.Errorf("module not registered: %s", name)
	}
	return m, nil
}

type Provisioner interface {
	Provision(ctx Context) error
}

type Validator interface {
	Validate() error
}

type CleanerUpper interface {
	Cleanup() error
}

var (
	modules   = make(map[string]ModuleInfo)
	modulesMu sync.RWMutex
)

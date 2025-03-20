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
	ID  string
	New func() Module
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

func GetModules(scope string) []ModuleInfo {
	modulesMu.RLock()
	defer modulesMu.RUnlock()

	scopeParts := strings.Split(scope, ".")

	// handle the special case of an empty scope, which
	// should match only the top-level modules
	if scope == "" {
		scopeParts = []string{}
	}

	var mods []ModuleInfo
iterateModules:
	for id, m := range modules {
		modParts := strings.Split(id, ".")

		// match only the next level of nesting
		if len(modParts) != len(scopeParts)+1 {
			continue
		}

		// specified parts must be exact matches
		for i := range scopeParts {
			if modParts[i] != scopeParts[i] {
				continue iterateModules
			}
		}

		mods = append(mods, m)
	}

	// make return value deterministic
	sort.Slice(mods, func(i, j int) bool {
		return mods[i].ID < mods[j].ID
	})

	return mods
}

func Modules() []string {
	modulesMu.RLock()
	defer modulesMu.RUnlock()

	names := make([]string, 0, len(modules))
	for name := range modules {
		names = append(names, name)
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

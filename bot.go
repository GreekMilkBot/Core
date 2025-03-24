package Core

import (
	"context"
	"fmt"
	"slices"
	"sync"
)

type Bot struct {
	locker    *sync.RWMutex
	instances []BotInstance
}

type BotInstance struct {
	ModuleInfo
	Mod    *Module
	Config map[string]any
}

func NewBot() *Bot {
	return &Bot{
		locker:    new(sync.RWMutex),
		instances: make([]BotInstance, 0),
	}
}

func (b *Bot) Add(name string, cfg map[string]any) error {
	return b.addOrSet(true, name, cfg)
}
func (b *Bot) addOrSet(create bool, name string, cfg map[string]any) error {
	b.locker.Lock()
	defer b.locker.Unlock()
	if cfg == nil {
		cfg = make(map[string]any)
	}
	module, err := GetModule(name)
	if err != nil {
		return err
	}
	var mod *BotInstance = nil
	if index := slices.IndexFunc(b.instances, func(x BotInstance) bool {
		return x.ID == name
	}); index != -1 {
		if !b.instances[index].Multi {
			if !create {
				return fmt.Errorf("bot instance `%s` already exists", name)
			} else {
				mod = &b.instances[index]
			}
		}
	}
	if mod == nil {
		if !create {
			return fmt.Errorf("bot instance `%s` not found", name)
		}
		m := module.New()
		mod = &BotInstance{
			ModuleInfo: module,
			Mod:        &m,
		}
		b.instances = append(b.instances, *mod)
	}
	mod.Config = cfg

	return nil
}

type BotProcessor interface {
	Processor(mods *[]BotInstance) error
}

func (b *Bot) Start(ctx context.Context) {}

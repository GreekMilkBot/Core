package Core

import (
	"code.d7z.net/d7z-team/go-variables"
	"context"
	"github.com/pkg/errors"
	"slices"
	"sync"
)

type Bot struct {
	locker    *sync.RWMutex
	instances []BotInstance
	ctx       Context
}

type BotInstance struct {
	ModuleInfo
	Mod    Module
	Config variables.Variables
	Booted bool
	closer []func() error
}

func (b *BotInstance) ctx(ctx Context) Context {
	ctx = ctx.WithValue("name", b.ID)
	ctx.Config = b.Config
	return ctx
}

func (b *BotInstance) init(ctx Context) error {
	if b.Booted {
		return nil
	}
	b.Booted = true
	if item, ok := b.Mod.(Provisioner); ok {
		if err := item.Provision(b.ctx(ctx)); err != nil {
			return err
		}
	}
	if item, ok := b.Mod.(Validator); ok {
		if err := item.Validate(); err != nil {
			return err
		}
	}
	if item, ok := b.Mod.(CleanerUpper); ok {
		b.closer = append(b.closer, item.Cleanup)
	}
	return nil
}

func NewBot(ctx context.Context) *Bot {
	r := &Bot{
		locker:    new(sync.RWMutex),
		instances: make([]BotInstance, 0),
		ctx:       NewContext(ctx),
	}
	for _, s := range Modules("system.") {
		if err := r.Add(s, nil); err != nil {
			panic(err)
		}
	}
	return r
}

func (b *Bot) Add(name string, cfg map[string]any) error {
	b.locker.Lock()
	defer b.locker.Unlock()
	if cfg == nil {
		cfg = make(map[string]any)
	}
	module, err := GetModule(name)
	if err != nil {
		return err
	}

	mod := BotInstance{
		ModuleInfo: module,
		Mod:        module.New(),
	}
	b.instances = append(b.instances, mod)
	if data, ok := b.ctx.Value(mod.ID).(*BotInstances); ok && data != nil {
		data.plugins[mod.ID] = append(data.plugins[mod.ID], &mod.Mod)
	} else {
		b.ctx = b.ctx.WithValue(mod.ID, &BotInstances{
			plugins: map[string][]*Module{
				mod.ID: {&mod.Mod},
			},
		})
	}
	mod.Config = cfg
	return b.afterHook()
}

func (b *Bot) Update(name string, cfg map[string]any, index int) error {
	b.locker.Lock()
	defer b.locker.Unlock()
	if cfg == nil {
		cfg = make(map[string]any)
	}
	var mod *BotInstance
	for _, i := range b.instances {
		if i.ID == name {
			if index == 0 {
				mod = &i
				break
			}
			index--
		}
	}
	if mod == nil {
		return errors.New("module not found")
	}
	mod.Config = cfg
	return b.afterHook()
}

func (b *Bot) afterHook() error {
	type hookItem struct {
		*BotInstance
		BotProcessor
	}
	hooks := make([]hookItem, 0)
	for _, instance := range b.instances {
		mod := instance.Mod
		if item, ok := mod.(BotProcessor); ok && item != nil {
			hooks = append(hooks, hookItem{
				BotInstance:  &instance,
				BotProcessor: item,
			})
		}
	}

	slices.SortFunc(hooks, func(i, j hookItem) int {
		var a, b int
		if item, ok := i.Mod.(Priority); ok {
			a = item.Priority()
		}
		if item, ok := j.Mod.(Priority); ok {
			b = item.Priority()
		}
		return a - b
	})
	for _, hook := range hooks {
		if err := hook.init(b.ctx); err != nil {
			return err
		}
		if err := hook.Processor(&b.ctx, &b.instances); err != nil {
			return err
		}
	}
	return nil

}

type BotProcessor interface {
	Processor(context *Context, mods *[]BotInstance) error
}

func (b *Bot) Start() {}
